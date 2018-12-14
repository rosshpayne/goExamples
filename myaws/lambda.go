package main
 
import (
        "context"
        "log"
        "time"
//        "github.com/aws/aws-lambda-go/lambda"
)
 
func LongRunningHandler(ctx context.Context) string {
        deadline, _ := ctx.Deadline()
        yy := make(chan time.Duration)
        for {
                select {
                case   yy<- time.Until(deadline).Truncate(100 * time.Millisecond):
                       // fmt.Printf("xx = %v",xx)
                        return "Finished before timing out."
                default:
                        log.Print("hello!")
                        time.Sleep(50 * time.Millisecond)
                }
        }
        close(yy)
}
 
func main() {
//       lambda.Start(LongRunningHandler)        
         LongRunningHandler(context.Background())
}
