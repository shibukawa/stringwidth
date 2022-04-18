package stringwidth

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		src string
		opt []Opt
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple words",
			args: args{
				src: "hello",
				opt: []Opt{{IsAmbiguousWide: false}},
			},
			want: 5,
		},
		{
			name: "East Asian Wide",
			args: args{
				src: "こんにちわ",
				opt: []Opt{{IsAmbiguousWide: false}},
			},
			want: 10,
		},
		{
			name: "East Asian Ambiguous (option false)",
			args: args{
				src: "¼½¾",
				opt: []Opt{{IsAmbiguousWide: false}},
			},
			want: 3,
		},
		{
			name: "East Asian Ambiguous (option true)",
			args: args{
				src: "¼½¾",
				opt: []Opt{{IsAmbiguousWide: true}},
			},
			want: 6,
		},
		{
			name: "Standard Emoji",
			args: args{
				src: "🐙",
				opt: []Opt{{IsAmbiguousWide: false}},
			},
			want: 2,
		},
		{
			name: "Emoji Sequence",
			args: args{
				src: "🏴󠁧󠁢󠁥󠁮󠁧󠁿", // 0x1F3F4 0xE0067 0xE0062 0xE0065 0xE006E 0xE0067 0xE007F
				opt: []Opt{{IsAmbiguousWide: true}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.src, tt.args.opt...); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
