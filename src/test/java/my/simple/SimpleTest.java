package my.simple;

import org.junit.Test;
import org.junit.experimental.categories.Category;

public class SimpleTest {

    @Test
    @Category({FirstCategory.class})
    public void runTest(){
        System.out.println("runTest");
    }

    @Test
    @Category({SecondCategory.class})
    public void runTest2(){
        System.out.println("runSecondTest");
    }

}
