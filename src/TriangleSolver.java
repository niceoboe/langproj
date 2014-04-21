import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.util.ArrayList;
import java.util.Scanner;

public class TriangleSolver {
	
	public TriangleSolver(String filename) throws FileNotFoundException {
		Scanner fin = new Scanner(new BufferedReader(new FileReader(filename))); //input file
		String line;
		triangle = new ArrayList<>();
		String[] strArray; //Array of numbers in the row as a string
		int[] valueArray; //int values of an strArray
		
		//Populates triangle with input from file
		while(fin.hasNextLine()){
			line = fin.nextLine();
			strArray = line.split(" ");
			valueArray = new int[strArray.length];
			for(int i = 0; i < valueArray.length; i++)
				valueArray[i] = Integer.parseInt(strArray[i]);
			triangle.add(valueArray);
		}
		fin.close();
	}
	
	private ArrayList<int[]> triangle;
	
	public static void main(String args[]){
		try{
			String filename = args[0];
			TriangleSolver tri = new TriangleSolver(filename);
			System.out.println(tri.maximum());
		}
		catch(FileNotFoundException e){
			System.out.println("Format: java TriangleSolver <file>");
		}
	}
	
	public int maximum(){
		//Starts at the second to last row and "collapses" the row below it
		for(int i = triangle.size() - 2; i >= 0; i--){
			for(int j = 0; j < triangle.get(i).length; j++){
				//Adds a child node to the parent node which will produce the largest sum
				if(triangle.get(i)[j] + triangle.get(i+1)[j] >= triangle.get(i)[j] + triangle.get(i+1)[j+1]){
					triangle.get(i)[j] += triangle.get(i+1)[j];
				}
				else{
					triangle.get(i)[j] += triangle.get(i+1)[j+1];
				}
			}
		}
		return triangle.get(0)[0];
	}
}
