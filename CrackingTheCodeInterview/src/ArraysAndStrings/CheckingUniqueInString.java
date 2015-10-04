package ArraysAndStrings;

public class CheckingUniqueInString {
	public static boolean checkUniqueBool(String s)
	{
		for(int i=0;i<s.length();i++)
		{
			for(int j=i+1;j<s.length();j++)
			{
				if(s.charAt(i)==s.charAt(j))
				{
					return false;
				}
			}
		}
		return true;
	}
	//This is with order n complexity
	public static String checkUnique(String s)
	{
		if(s.length()==1)
		{
			return s;
		}
		 String character= s.substring(0,1);
		String newString=checkUnique(s.substring(1));
		{
			if(newString.contains(character)){
				return newString;
				
			}
			else
			{
				return newString+character;
			
			}
			
		}
	}
	public static void main(String[] args) {
		// TODO Auto-generated method stub

		String first="Geetika";
		String second=checkUnique(first);
		if(second.length()==first.length())
		{
			System.out.println("True");
		}
		else
		{
			System.out.println("False");
		}
		System.out.print(checkUniqueBool(first));
	}

}