package Mapping;

public class HashMapUse {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Map m = new Map();
		m.add("abcff", "def");
		System.out.println("abcff".hashCode()+" "+"abcd".hashCode()+" "+"ghi".hashCode());
		m.add("ghi", "xyz");
		System.out.println(m.getValue("abcff"));
		//m.add("geetika", "good");
		m.add("abcd", "yes");
		//m.getValue("geetika");
	}

}
