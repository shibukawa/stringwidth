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
				opt: []Opt{},
			},
			want: 5,
		},
		{
			name: "East Asian Wide",
			args: args{
				src: "こんにちわ",
				opt: []Opt{},
			},
			want: 10,
		},
		{
			name: "East Asian Ambiguous (option default)",
			args: args{
				src: "¼½¾",
				opt: []Opt{},
			},
			want: 3,
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
				opt: []Opt{},
			},
			want: 2,
		},
		{
			name: "Emoji Sequence",
			args: args{
				src: "🏴󠁧󠁢󠁥󠁮󠁧󠁿", // 0x1F3F4 0xE0067 0xE0062 0xE0065 0xE006E 0xE0067 0xE007F
				opt: []Opt{},
			},
			want: 2,
		},
		{
			name: "ANSI escape code",
			args: args{
				src: "\x1b[38;5;140m foo\x1b[0m bar",
				opt: []Opt{},
			},
			want: 8,
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
