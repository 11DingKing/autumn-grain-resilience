package health
import "testing"
import "context"
func TestTask021(t *testing.T){if Check(context.Background(),nil).Healthy{t.Fatal("nil database reported healthy")}}