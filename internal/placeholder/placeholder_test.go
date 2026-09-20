package placeholder

import "testing"

func TestExpandUsesValuesAndDefaults(t *testing.T) {
	got, err := Expand("docker run --name ${NAME:=web} ${IMAGE}", map[string]string{
		"IMAGE": "nginx",
	})
	if err != nil {
		t.Fatal(err)
	}
	want := "docker run --name web nginx"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestExpandReportsMissingValue(t *testing.T) {
	_, err := Expand("echo ${NAME}", map[string]string{})
	if err == nil {
		t.Fatal("expected a missing placeholder value to fail")
	}
}
