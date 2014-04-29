use std::path::Path;
use std::io::fs::File;
use std::io::buffered::BufferedReader;



fn main() {
	
	// Specifies the Path for the Triangle text file
	let triangle_file_path = Path::new("triangle.txt".as_slice());


	// Opens the file with the specified path. 
	// Note the use of the "&" (address of) operator
	let triangle_file = File::open(&triangle_file_path);

	// Opens new BufferedReader for this file, allowing us to read it line-by-line
	// Note the use of the "mut" keyword, which declares this variable as "mutable" and allows it to be read in the future
	let mut triangle_file_reader = BufferedReader::new(triangle_file);

	// The triangle itself
	// This syntax instantiates the triangle as a 2D array of ints
	let mut triangle: ~[~[int]] = ~[];
	
	let mut rows_loaded = 0;


	// Iterates through each line in the BufferedReader
	for line in triangle_file_reader.lines() {
		// A temporary 1D array of ints that represents a row of the triangle
		let mut triangle_row: ~[int] = ~[];

		for i in range (0, 100) {

			if i <= rows_loaded {
				match from_str::<int>(line.slice((i*3) as uint, ((i*3) + 2) as uint)) {
					Some(x) => triangle_row.push(x),
					None => fail!("Conversion failed!")
				};
			}
		}
		triangle.push(triangle_row);
		rows_loaded += 1;
	}
}

fn prepend<T>(xs: List<T>, value: T) -> List<T> {
	Cons(value, ~xs)
}

enum List<T> {
	Cons(T, ~List<T>),
	Nil
}