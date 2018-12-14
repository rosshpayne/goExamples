package main

import (
       "fmt"
       "time"
       )

type point struct {
             x, y int
                 }

func (p *point) sum() int {
            return (*p).x + (*p).y
             }
func (p *point) string() string {
            return fmt.Sprintf("A Point { %d, %d %d}",(*p).x, (*p).y, (*p).sum())
             }

type notpoint struct {
             x, y int
                 }

func (p *notpoint) sum() int {
            return p.x - p.y         // go will auto dereference p i.e (*p) 
             }
func (p *notpoint) string() string {
            return fmt.Sprintf("Not a point { %d, %d %d}",p.x, p.y, p.sum())
            }

type xpoint struct {
             x, y int
                 }

func (p xpoint) sum() int {
            return p.x + p.y         // go will auto dereference p i.e (*p) 
             }
func (p xpoint) string() string {
            return fmt.Sprintf("This is an x point{ %d, %d %d }",p.x, p.y,p.sum())
            }

type ipoint interface {
            sum() int
            string() string 
                    }

type circle  struct {
            ipoint               // interface 
            radius float32
                   }

func main () {

             const maxUint64 = 1<<31 - 1

             var aLongTimeAgo = time.Unix(1, 0)

             acircle := circle{&point{1,2},2.5}
             ncircle := circle{&notpoint{1,2},10}
             xcircle := circle{xpoint{3,4},14}
             mypoint := point{4,5}
             fmt.Println("mypoint: ",mypoint.string());   // go automatically adds &mypoint.string()
             fmt.Println("mypoint: ",(&mypoint).string());   // go automatically adds &mypoint.string()

             fmt.Println(acircle.string())
             fmt.Println(ncircle.string())
             fmt.Println(xcircle.string())

             fmt.Printf("%b %d\n",maxUint64,maxUint64)
             fmt.Printf("%s\n",aLongTimeAgo)
             fmt.Printf("%s\n",time.Now())
             
             xx:=time.Time{}
             fmt.Printf("%T\n",xx)
             y,m,d:=xx.Date()
             fmt.Println(y, m, d)
             
	             crlf   := []byte("stu")
                     for _,r := range crlf {
                          fmt.Printf("%b %c %d %x\n",r,r,r,r)
                     }
             //
             
             billion := float64(1000000000)
             b:=billion
             l:=1e-06
             fmt.Println(billion)
             fmt.Printf("%f\n",l)
             for i := 0;i<1000000;i++ { b+=l }
             fmt.Printf("%9.8f\n",b-billion)
             //
}
