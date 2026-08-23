package notifications
import("context";"testing")
func TestTask025(t *testing.T){sender:=&MemorySender{Failures:2};d:=NewDispatcher();d.MaxAttempts=1;d.Senders[SMS]=sender;if e:=d.Dispatch(context.Background(),Message{ID:"m",Recipient:"farmer",Body:"warning",Channel:SMS});e==nil{t.Fatal("delivery failure hidden")};if _,ok:=d.Delivered("m");ok{t.Fatal("failed notification recorded as delivered")}}