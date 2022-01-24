package recursion

import (
	"reflect"
	"testing"
)

func TestItem_IsDoll(t *testing.T) {
	type fields struct {
		ID    int
		Type  string
		Child *Item
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "test01",
			fields: fields{ID: 0, Type: "doll", Child: nil},
			want:   true,
		},
		{
			name: "test02",
			fields: fields{
				ID:    1,
				Type:  "diamond",
				Child: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := &Item{
				ID:    tt.fields.ID,
				Type:  tt.fields.Type,
				Child: tt.fields.Child,
			}
			if got := it.IsDoll(); got != tt.want {
				t.Errorf("IsDoll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDiamond(t *testing.T) {
	type args struct {
		item Item
	}
	tests := []struct {
		name string
		args args
		want Item
	}{
		{
			name: "test01",
			args: args{Item{ID: 0, Type: "doll", Child: &Item{ID: 1, Type: "diamond", Child: nil}}},
			want: Item{ID: 1, Type: "diamond", Child: nil},
		},
		{
			name: "test02",
			args: args{
				Item{ID: 0, Type: "doll",
					Child: &Item{ID: 1, Type: "doll",
						Child: &Item{ID: 2, Type: "diamond",
							Child: nil,
						}}}},
			want: Item{ID: 2, Type: "diamond", Child: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiamond(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiamond() = %v, want %v", got, tt.want)
			}
		})
	}
}
