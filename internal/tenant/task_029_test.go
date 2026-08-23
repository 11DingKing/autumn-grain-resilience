package tenant
import "testing"
func TestTask029(t *testing.T){r:=NewRegistry();_ = r.Set("operator",Scope{RegionIDs:map[string]bool{"north":true}});scope,_:=r.Get("operator");scope.RegionIDs["south"]=true;stored,_:=r.Get("operator");if stored.Allows("south"){t.Fatal("returned scope mutated authorization registry")}}