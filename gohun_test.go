package gohun

import (
	"io/ioutil"
	"os"
	"testing"
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
	expected := []string{"carol", "valor", "color", "cal or", "cal-or",
		"caloric", "calorie", "chloral", "Carlo", "Calgary", "Caloocan"}
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

func TestIsCorrect(t *testing.T) {
	g, err := getGohun()
	if err != nil {
		t.Error("Failed to initialize Gohun struct:" + err.Error())
	} else {
		w := "calor"
		if g.IsCorrect(w) {
			t.Errorf("IsCorrect(\"%s\") failed, it returned true when it should have returned false.", w)
		}
		w = "color"
		if !g.IsCorrect(w) {
			t.Errorf("IsCorrect(\"%s\") failed, it returned false when it should have returned true.", w)
		}
	}
}

func TestAddDictionary(t *testing.T) {
	expected := []string{"color", "co lour", "co-lour", "col our", "col-our", "cornflour",
		"Colo", "contour", "courtly", "Colbert", "colonize"}
	g, err := getGohun()
	if err != nil {
		t.Errorf("Failed to initialize gohun object:" + err.Error())
	} else {
		w := "colour"
		b, n, sugg := g.CheckSuggestions(w)
		if b || n != 11 || !compareSlices(sugg, expected) {
			t.Errorf("CheckSuggestions(\"%s\") pre add failed, it returned: %t, %d, %+v, and expected: %t, %d, %+v",
				w, b, n, sugg, false, 11, expected)
			return
		}
		file, err := os.Open("./include/dictionaries/en_CA.dic")
		if err != nil {
			t.Errorf("Openfile pre add failed:" + err.Error())
			return
		}
		dic, e := ioutil.ReadAll(file)
		file.Close()
		if e != nil {
			t.Errorf("Failed to read string from file pre add:" + err.Error())
			return
		}
		e = g.AddDictionary(dic)
		if e != nil {
			t.Errorf("AddDictionary(string) failed to add dictionary: " + e.Error())
			return
		}
		b2, _, _ := g.CheckSuggestions(w)
		if !b2 {
			t.Errorf("CheckSuggestions(\"%s\") post add failed. Was expecting the word \"%s\" to be correct.", w, w)
		}
	}
}

func TestAddWord(t *testing.T) {
	expected := []string{"color", "co lour", "co-lour", "col our", "col-our", "cornflour",
		"Colo", "contour", "courtly", "Colbert", "colonize"}
	g, err := getGohun()
	if err != nil {
		t.Errorf("Failed to initialize gohun object: " + err.Error())
	} else {
		w := "colour"
		b, n, sugg := g.CheckSuggestions(w)
		if b || n != 11 || !compareSlices(sugg, expected) {
			t.Errorf("CheckSuggestions(\"%s\") pre add failed, it returned: %t, %d, %+v, and expected: %t, %d, %+v",
				w, b, n, sugg, false, 11, expected)
			return
		}
		b = g.AddWord(w)
		if !b {
			t.Errorf("AddWord(\"%s\") failed.", w)
			return
		}
		b, _, _ = g.CheckSuggestions(w)
		if !b {
			t.Errorf("\"%s\" returned as being an incorrect word despite being added to the gohun object.")
		}
	}
}

func TestRemoveWord(t *testing.T) {
	expected := []string{"colon", "dolor", "col or", "col-or", "colored", "recolor",
		"colorful", "Colorado", "Colon", "colorize"}
	g, err := getGohun()
	if err != nil {
		t.Errorf("Failed to initialize gohun object: " + err.Error())
	} else {
		w := "color"
		b, _, _ := g.CheckSuggestions(w)
		if !b {
			t.Errorf("CheckSuggestions(\"%s\") pre remove failed, it invalidate \"%s\" as an incorrect word, though it should be in the dictionary", w, w)
			return
		}
		b = g.RemoveWord(w)
		if !b {
			t.Errorf("RemoveWord(\"%s\") failed.", w)
			return
		}
		b2, n, sugg := g.CheckSuggestions(w)
		if b2 || n != 10 || !compareSlices(sugg, expected) {
			t.Errorf("CheckSuggestions(\"%s\") post remove failed, it returned: %t, %d, %+v, and expected: %t, %d, %+v",
				w, b2, n, sugg, false, 10, expected)
		}
	}
}

func TestStem(t *testing.T) {
	expected := []string{"telling", "tell"}
	g, err := getGohun()
	if err != nil {
		t.Errorf("Failed to initialize gohun object: " + err.Error())
	} else {
		w := "telling"
		n, sugg := g.Stem(w)
		if n != 2 || !compareSlices(sugg, expected) {
			t.Errorf("Stem(\"%s\") failed. Expected: %d, %v; got: %d, %v", w, 2, expected, n, sugg)
		}
	}
}

func TestGenerate(t *testing.T) {
	expected := []string{"told"}
	nexp := 1
	g, err := getGohun()
	if err != nil {
		t.Errorf("Failed to initialize gohun object: " + err.Error())
	} else {
		w1 := "telling"
		w2 := "ran"
		n, sugg := g.Generate(w1, w2)
		if n != nexp || !compareSlices(sugg, expected) {
			t.Errorf("Generate(\"%s\",\"%s\") failed. Expected: %d, %v; got: %d, %v", w1, w2, nexp, expected, n, sugg)
		}
	}
}

func TestAnalyze(t *testing.T) {
	expected := []string{" st:telling ts:0", " st:tell ts:0 al:told is:Vg"}
	nexp := 2
	g, err := getGohun()
	if err != nil {
		t.Errorf("Failed to initialize gohun object: " + err.Error())
	} else {
		w := "telling"
		n, sugg := g.Analyze(w)
		if n != nexp || !compareSlices(sugg, expected) {
			t.Errorf("Analyze(\"%s\") failed. Expected: %d, %v; got: %d, %v", w, nexp, expected, n, sugg)
		}
	}
}
