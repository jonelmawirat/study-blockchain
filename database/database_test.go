package database 

import (
    "os"
	"testing"
    "io/ioutil"
)


func TestAccount(t *testing.T) {

    var test_account Account
    test_account = Account("Jonel")    
    want := "Jonel"
    if test_account != Account(want) {
        t.Errorf("Want: %s Got: %s\n",want,test_account)
    }
    
}

func TestGenesis(t *testing.T) {

    type testCase struct {
        name        string
        got         uint 
        want        uint 
    }

    tempFile,err := ioutil.TempFile("","test_genesis_*.json")
    if err != nil {
        t.Error(err)
    } 

    defer os.Remove(tempFile.Name())

    if _,err := tempFile.Write([]byte(`{"balances": {"1":1,"2":2}}`)); err != nil {
        t.Error(err)
    }

    tempFile.Close()

    genesis_data,err := loadGenesis(tempFile.Name())
    if err != nil {
        t.Error(err)
    }

    testCases := []testCase{
        {name: "Get Account Balance",got: genesis_data.Balances["1"],want:1},
        {name: "Get Account Balance",got: genesis_data.Balances["2"],want:2},
    }

    for _,tc := range testCases {
        if tc.got != tc.want {
            t.Errorf("%s - Want: %d Got: %d\n",tc.name,tc.want,tc.got)
        }
    }

    
     

}
