package lease
import("context";"testing";"time")
func TestTask015(t *testing.T){m:=New();_,_ = m.Acquire(context.Background(),"x","a",time.Minute);if e:=m.Renew(context.Background(),"x","b",time.Minute);e==nil{t.Fatal("foreign renew")}}