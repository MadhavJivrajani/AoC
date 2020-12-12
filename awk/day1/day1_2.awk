# Run as: awk -f day1_2.awk input.txt input.txt

NR == FNR { map[$1]; next }

1 {
	for(e in map) {
		if(e == $1) { continue }
		if((2020 - $1 - e) in map) { print $1 * e * (2020 - $1 - e); exit }
	}
}
