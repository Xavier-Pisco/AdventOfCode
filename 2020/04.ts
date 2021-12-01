import * as files from 'fs';

let input: string[] = files.readFileSync('04.txt').toString().split('\n\n');
let users: string[][] = [];

for (let i = 0; i < input.length; i++) {
	let part1 = input[i].split(new RegExp(' |\n'));
	users.push(part1);
}
let validPassports = 0;
for (let i = 0; i < users.length; i++) {
	let userFields: string[] = [];
	let validFields1 = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid', 'cid'];
	let validFields2 = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'];
	validFields1.sort();
	validFields2.sort();

	for (let j = 0; j < users[i].length; j++) {
		let fields = users[i][j].split(':');

		if (fields[0] === 'byr') {
			if (Number(fields[1]) <= 2002 && Number(fields[1]) >= 1920)
				userFields.push(fields[0]);
		} else if (fields[0] === 'iyr') {
			if (Number(fields[1]) >= 2010 && Number(fields[1]) <= 2020)
				userFields.push(fields[0]);
		} else if (fields[0] === 'eyr') {
			if (Number(fields[1]) >= 2020 && Number(fields[1]) <= 2030)
				userFields.push(fields[0]);
		} else if (fields[0] === 'hgt') {
			if (fields[1].length === 4 && fields[1][2] === 'i' && fields[1][3] === 'n') {
				let height = Number(fields[1][0]) * 10 + Number(fields[1][1]);
				if (height <= 76 || height >= 59) {
					userFields.push(fields[0]);

				}
			} else if (fields[1].length === 5 && fields[1][3] === 'c' && fields[1][4] === 'm') {
				let height = Number(fields[1][0]) * 100 + Number(fields[1][1]) * 10 + Number(fields[1][2]);
				console.log(height);

				if (height >= 150 || height <= 193) {
					userFields.push(fields[0]);
				}
			}
		} else if (fields[0] === 'hcl') {
			let validChars = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'];
			let numberValidChars = 0;
			if (fields[1][0] === '#') {
				for (let x = 1; x < fields[1].length; x++) {
					if (validChars.includes(fields[1][x])) {
						numberValidChars += 1;
					}
				}
			}
			if (numberValidChars === 6) {
				userFields.push(fields[0]);
			}
		} else if (fields[0] === 'ecl') {
			let validStrings = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'];
			if (validStrings.includes(fields[1])){
				userFields.push(fields[0]);
			}
		} else if (fields[0] === 'pid') {
			if (fields[1].length === 9) {
				for (let x = 0; x < fields[1].length; x++) {
					if (fields[1][x] >= '0' && fields[1][x] <= '9'){
						userFields.push(fields[0]);
						break;
					}
				}
			}
		} else if (fields[0] === 'cid') {
			fields.push(fields[0]);
		}
	}
	userFields.sort();
	console.log(userFields);

	if (userFields.toString() === validFields1.toString() || userFields.toString() === validFields2.toString()) {
		validPassports += 1;
	}
}
console.log('2nd Part: ' + validPassports);

