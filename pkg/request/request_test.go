package request

import "testing"

func Test_validateMutualExclusiveFields(t *testing.T) {
	type args struct {
		fields map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "all fields empty",
			args: args{
				fields: map[string]string{
					"txt": "",
					"url": "",
					"doc": "",
				},
			},
			wantErr: true,
		},
		{
			name: "all fields set",
			args: args{
				fields: map[string]string{
					"txt": "text",
					"url": "url",
					"doc": "doc",
				},
			},
			wantErr: true,
		},
		{
			name: "one field set",
			args: args{
				fields: map[string]string{
					"txt": "text",
					"url": "",
					"doc": "",
				},
			},
			wantErr: false,
		},
		{
			name: "two fields set",
			args: args{
				fields: map[string]string{
					"txt": "text",
					"url": "url",
					"doc": "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateMutualExclusiveFields(tt.args.fields); (err != nil) != tt.wantErr {
				t.Errorf("validateMutualExclusiveFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
