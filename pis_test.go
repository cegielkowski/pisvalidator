package pisvalidatior

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		pis     string
		wantErr bool
	}{
		{
			pis:     "71115848273",
			wantErr: false,
		},
		{
			pis:     "477.11617.27-5",
			wantErr: false,
		},
		{
			pis:     "475.12417.27-5",
			wantErr: true,
		},
		{
			pis:     "47311117177",
			wantErr: true,
		},
		{
			pis:     "4731111717",
			wantErr: true,
		},
		{
			pis:     "47a11117177",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.pis, func(t *testing.T) {
			if _, err := ValidatePis(tt.pis); (err != nil) != tt.wantErr {
				t.Errorf("Error on validate = %v, wantErr %v, pis %v", err, tt.wantErr, tt.pis)
			}
		})
	}
}

func TestClean1(t *testing.T) {
	s := clean("111.222.333-99")

	if s != "11122233399" {
		t.Errorf("Invalid result: %s", s)
	}
}

func TestClean2(t *testing.T) {
	s := clean("10.963.268/0001-82")

	if s != "10963268000182" {
		t.Errorf("Invalid result: %s", s)
	}
}

func TestDigit1(t *testing.T) {
	s := verifyDigit(11)

	if s != 0 {
		t.Errorf("Invalid result: %d", s)
	}
}

func TestDigit2(t *testing.T) {
	s := verifyDigit(9)

	if s != 9 {
		t.Errorf("Invalid result: %d", s)
	}
}
