package db

import (
	"backend/config"
	"backend/log"

	"context"
	"encoding/json"
	"errors"

	"github.com/dgraph-io/dgraph/client"
	"github.com/dgraph-io/dgraph/protos/api"
	"github.com/segmentio/bellows"
	"google.golang.org/grpc"
)

var cl *client.Dgraph

const (
	ID = "uid"
)

func create(entity interface{}) (string, error) {
	flatMap := bellows.Flatten(entity)
	delete(flatMap, ID)

	txn := cl.NewTxn()
	defer txn.Discard(context.Background())

	log.Get().WithField("entity", flatMap).Debug("Delta going in")

	bs, err := json.Marshal(flatMap)
	if err != nil {
		log.Get().WithError(err).Error("While marshalling mutation")
		return "", err
	}

	mu := &api.Mutation{
		SetJson: bs,
	}
	as, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		log.Get().WithError(err).Error("While saving delta")
		return "", err
	}

	if err := txn.Commit(context.Background()); err != nil {
		log.Get().WithError(err).Error("While commiting mutation")
		return "", err
	}

	for _, uid := range as.GetUids() {
		log.Get().WithField("uid", uid).Debug("Created uid")
		return uid, nil
	}

	return "", errors.New("No entity created")
}

func update(id string, entity interface{}) error {
	flatMap := bellows.Flatten(entity)
	flatMap[ID] = id

	log.Get().WithField("Delta", flatMap).Debug("Delta update going in")

	txn := cl.NewTxn()
	defer txn.Discard(context.Background())

	bs, err := json.Marshal(flatMap)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		SetJson: bs,
	}
	as, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		log.Get().WithError(err).Error("While marshalling mutation")
		return err
	}

	if err := txn.Commit(context.Background()); err != nil {
		log.Get().WithError(err).Error("While commiting changes")

		return err
	}

	for _, uid := range as.GetUids() {
		if id != uid {
			log.Get().WithField("Original id", id).WithField("New id", uid).Error("While updating predicates, id do not match")
			return errors.New("Returned and supplied UIDs do not match")
		}

		return nil
	}

	return nil
}

func addEdge(sub, pred, obj string, vm ...map[string][]byte) error {
	txn := cl.NewTxn()
	defer txn.Discard(context.Background())
	log.Get().WithField("Subject", sub).WithField("Predicate", pred).WithField("Object", obj).WithField("Facets", vm).Debug("Adding edge")

	mu := &api.Mutation{}
	if len(vm) > 0 {
		var facs []*api.Facet
		for key, val := range vm[0] {
			f := api.Facet{Key: key, Value: val}
			facs = append(facs, &f)
		}
		mu.Set = append(mu.Set, &api.NQuad{
			Subject:   sub,
			Predicate: pred,
			ObjectId:  obj,
			Facets:    facs,
		})
	} else {
		mu.Set = append(mu.Set, &api.NQuad{
			Subject:   sub,
			Predicate: pred,
			ObjectId:  obj,
		})
	}
	if _, err := txn.Mutate(context.Background(), mu); err != nil {
		log.Get().WithError(err).Error("While adding predicates")
		return err
	}

	return txn.Commit(context.Background())
}

func deleteAllPreds(id string, pred ...string) error {
	txn := cl.NewTxn()
	defer txn.Discard(context.Background())

	log.Get().WithField("Subject", id).WithField("Predicates", pred).Debug("Deleting predicates")
	mu := &api.Mutation{}
	client.DeleteEdges(mu, id, pred...)

	if _, err := txn.Mutate(context.Background(), mu); err != nil {
		log.Get().WithError(err).Error("While deleting predicates")
		return err
	}

	return txn.Commit(context.Background())
}

func deleteEdges(id string, pred string, edges ...string) error {
	txn := cl.NewTxn()
	defer txn.Discard(context.Background())
	log.Get().WithField("Subject", id).WithField("Predicate", pred).WithField("Edges", edges).Debug("Deleting edges")

	mu := &api.Mutation{}
	for _, edge := range edges {
		mu.Del = append(mu.Del, &api.NQuad{
			Subject:   id,
			Predicate: pred,
			ObjectId:  edge,
		})
	}
	if _, err := txn.Mutate(context.Background(), mu); err != nil {
		log.Get().WithError(err).Error("While deleting edges")
		return err
	}

	return txn.Commit(context.Background())
}

func RawSearch(q string, varMap ...map[string]string) ([]byte, error) {
	log.Get().WithField("query", q).WithField("vars", varMap).Debug("Query going to dgraph")

	var res *api.Response
	var err error

	if len(varMap) < 1 {
		res, err = cl.NewTxn().Query(context.Background(), q)
	} else {
		res, err = cl.NewTxn().QueryWithVars(context.Background(), q, varMap[0])
	}

	if err != nil {
		log.Get().WithError(err).Error("While doing search")
		return nil, err
	}

	log.Get().WithField("result", string(res.Json)).Debug("Response from dgraph")

	return res.Json, nil
}

const max = 10 * 1024 * 1024

func Init(cfg *config.Config) (*grpc.ClientConn, error) {
	log.Get().Info("Connecting to dgraph...")
	conn, err := grpc.Dial(cfg.Dgraph, grpc.WithInsecure(), grpc.WithMaxMsgSize(max))
	if err != nil {
		return nil, err
	}

	dc := api.NewDgraphClient(conn)
	cl = client.NewDgraphClient(dc)

	return conn, err
}

func Get() *client.Dgraph {
	return cl
}

