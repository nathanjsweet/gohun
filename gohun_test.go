package gohun

import (
	"testing"
	"os"
	"io/ioutil"
)

func getGohun() (*Gohun, error) {
	file, err := os.Open("./include/dictionaries/en_US.aff")
	if err != nil {
		return nil, err
	}
	aff, e := ioutil.ReadAll(file)
	if e != nil {
		return nil, e
	}
	file.Close()
	file, err = os.Open("./include/dictionaries/en_US.dic")
	if err != nil {
		return nil, err
	}
	dic, e2 := ioutil.ReadAll(file)
	file.Close()
	if e2 != nil {
		return nil, e2
	}
	return NewGohun(aff, dic), nil
}

func compareSlices(f, s []string) bool {
	l1, l2 := len(f), len(s)
	if l1 != l2 {
		return false
	} else {
		for i := 0; i < l1; i++ {
			if f[i] != s[i] {
				return false
			}
		}
	}
	return true
}

func TestCheckSuggestions(t *testing.T) {
	expected := []string {"carol","valor","color","cal or","cal-or",
		"caloric","calorie","chloral","Carlo","Calgary","Caloocan"}
	g, err := getGohun()
	if err != nil {
		t.Error("Failed to initialize Gohun struct:" + err.Error())
	} else {
		w := "calor"
		b, n, sugg := g.CheckSuggestions(w)
		if b || n != 11 || !compareSlices(sugg, expected) {
			t.Errorf("CheckSuggestions(\"%s\") failed, it returned: %t, %d, %+v, and expected: %t, %d, %+v",
				w, b, n, sugg, false, 11, expected)
		}
	}
}

func TestAddDictionary(t *testing.T) {
	expected := []string {"color", "co lour", "co-lour", "col our", "col-our", "cornflour",
		"Colo", "contour", "courtly", "Colbert", "colonize"}
	g, err := getGohun()
	if err != nil {
		t.Error("Failed to initialize Gohun struct:" + err.Error())
	} else {
		w := "colour"
		b, n, sugg := g.CheckSuggestions(w)
		if b || n != 11 || !compareSlices(sugg, expected) {
			t.Errorf("AddDictionaryTest CheckSuggestions(\"%s\") pre add failed, it's  it returned: %t, %d, %+v, and expected: %t, %d, %+v",
				w, b, n, sugg, false, 11, expected)
			return
		}
		file, err := os.Open("./include/dictionaries/en_CA.dic")
		if err != nil {
			t.Errorf("AddDictionaryTest Openfile pre add failed:" + err.Error())
			return
		}
		dic, e := ioutil.ReadAll(file)
		file.Close()
		if e != nil {
			t.Errorf("AddDictionaryTest read string from file pre add failed:" + err.Error())
			return
		}
		b = g.AddDictionary(dic);
		if !b {
			t.Errorf("AddDictionaryTest AddDictionary(string) failed to add dictionary.")
			return
		}
		b2, _, _ := g.CheckSuggestions(w)
		if !b2 {
			t.Errorf("AddDictionaryTest CheckSuggestions(\"%s\") post add failed. Was expecting the word \"%s\" to be correct.", w, w)
		}
	}
}
