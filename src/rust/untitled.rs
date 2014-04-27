use std::io::File;

// Imports file-related methods
use std::io::BufferedReader;


fn main() {
	
	// Specifies the Path for the Triangle text file
	let triangle_file_path = Path::new("triangle.txt".as_slice());


	// Opens the file with the specified path. 
	// Note the use of the "&" (address of) operator
	// Note the use of the "mut" keyword, which declares this variable as "mutable" and allows it to be read in the future
	let mut triangle_file = File::open(&triangle_file_path);

	// Opens new BufferedReader for this file, allowing us to read it line-by-line
	let mut triangle_file_reader = BufferedReader::new(triangle_file);

	// Read the entire file into a variable of type String
	// let triangle_file_lines = triangle_file_reader.lines().map(|x| x.unwrap()).collect();

	let triangle_lines: ~[~str] = triangle_file_reader.lines().map(|x| x.unwrap()).collect();
	
	for line in triangle_lines.iter() {
		print!("{}", line);
	}
}