package geom

import (
	"reflect"
	"testing"
)

func TestSimplify(t *testing.T) {
	type simplifyTest struct {
		input     LineString
		output    LineString
		tolerance float64
	}

	var line = simplifyTest{
		input: LineString{
			Point{153.52, 928.49},
			Point{240.79, 988.95},
			Point{323.34, 1014.40},
			Point{404.41, 1020.08},
			Point{475.60, 981.17},
			Point{497.37, 921.45},
			Point{546.26, 903.57},
			Point{598.10, 907.57},
			Point{655.31, 941.11},
			Point{679.28, 1004.20},
			Point{630.91, 1052.36},
			Point{581.17, 1029.23},
		},
		output: LineString{
			Point{153.52, 928.49},
			Point{404.41, 1020.08},
			Point{546.26, 903.57},
			Point{655.31, 941.11},
			Point{679.28, 1004.20},
			Point{630.91, 1052.36},
			Point{581.17, 1029.23},
		},
		tolerance: 30.0,
	}
	var spiral = simplifyTest{
		input: LineString{
			Point{70.57, 609.01},
			Point{102.21, 618.89},
			Point{125.19, 635.79},
			Point{133.07, 659.34},
			Point{134.86, 688.40},
			Point{121.04, 709.80},
			Point{104.15, 726.70},
			Point{80.45, 731.71},
			Point{56.40, 729.34},
			Point{37.86, 714.81},
			Point{22.83, 692.69},
			Point{23.19, 669.21},
			Point{33.42, 648.38},
			Point{49.74, 635.79},
			Point{84.03, 628.63},
			Point{115.31, 645.31},
			Point{118.75, 681.96},
			Point{109.94, 704.43},
			Point{84.39, 715.17},
			Point{60.20, 716.24},
			Point{42.37, 703.14},
			Point{34.64, 675.16},
			Point{46.31, 658.05},
			Point{69.50, 645.16},
			Point{85.68, 651.96},
			Point{98.78, 669.93},
			Point{92.84, 691.98},
			Point{68.07, 699.21},
			Point{72.58, 676.59},
		},
		output: LineString{
			Point{70.57, 609.01},
			Point{133.07, 659.34},
			Point{104.15, 726.70},
			Point{37.86, 714.81},
			Point{23.19, 669.21},
			Point{84.03, 628.63},
			Point{118.75, 681.96},
			Point{60.20, 716.24},
			Point{46.31, 658.05},
			Point{98.78, 669.93},
			Point{72.58, 676.59},
		},
		tolerance: 30.0,
	}

	for _, test := range []simplifyTest{line, spiral} {
		o, err := Simplify(test.input, test.tolerance)
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(o, test.output) {
			t.Errorf("%v should equal %v.", o, test.output)
			t.Logf("%v should equal %v.", o, test.output)
		}
	}
}

func TestSimplifyInfiniteLoop1(t *testing.T) {
	geometry := Polygon{[]Point{Point{X: -871773.1638742175, Y: 497165.8489278648}, Point{X: -871786.058897913, Y: 497051.36283788737}, Point{X: -871812.6300702088, Y: 496843.16848807875}, Point{X: -871817.2594889546, Y: 496804.2641064115}, Point{X: -871828.8518461098, Y: 496707.8440110069}, Point{X: -871852.4900296698, Y: 496511.2579906229}, Point{X: -871868.0439736051, Y: 496404.69711395353}, Point{X: -871974.6604566738, Y: 496416.7107209433}, Point{X: -872080.2841608367, Y: 496430.28268054314}, Point{X: -872185.3395748375, Y: 496443.8991893595}, Point{X: -872291.2272915924, Y: 496457.2826704979}, Point{X: -872396.018682872, Y: 496470.4231297448}, Point{X: -872499.7279522702, Y: 496482.53964638617}, Point{X: -872578.9629476898, Y: 496491.58048275113}, Point{X: -872587.0595366888, Y: 496492.4496874893}, Point{X: -872593.1615890632, Y: 496493.189034136}, Point{X: -872651.0113261654, Y: 496500.8693724377}, Point{X: -872657.2059697689, Y: 496501.50819449686}, Point{X: -872664.1138388801, Y: 496502.2334463196}, Point{X: -872653.2429897713, Y: 496583.9830456376}, Point{X: -872641.9582658032, Y: 496669.1483742343}, Point{X: -872631.8505936791, Y: 496753.2266395176}, Point{X: -872611.5133181036, Y: 496923.7164973216}, Point{X: -872598.2396354877, Y: 497035.24964090344}, Point{X: -872585.005110706, Y: 497147.1231088713}, Point{X: -872571.3476570707, Y: 497259.16910962015}, Point{X: -872463.8327270573, Y: 497246.7009226391}, Point{X: -872354.5218117528, Y: 497233.79319054075}, Point{X: -872247.2050374286, Y: 497221.01688759495}, Point{X: -872130.8947989381, Y: 497206.8175871102}, Point{X: -872029.1911243061, Y: 497194.16539798304}, Point{X: -871936.3613199592, Y: 497183.25994469505}, Point{X: -871924.8707343122, Y: 497182.53941812366}, Point{X: -871878.9516291074, Y: 497176.64415429346}, Point{X: -871773.1638742175, Y: 497165.8489278648}, Point{X: -999943.8377197147, Y: 242782.39624921232}, Point{X: -999947.5075495535, Y: 242738.22203087248}, Point{X: -999958.0564939477, Y: 242680.11889008153}}}
	_, err := Simplify(geometry, 100.)
	if err != nil {
		t.Fatal(err)
	}
}
