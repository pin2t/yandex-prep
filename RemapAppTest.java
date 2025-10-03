import org.junit.jupiter.api.*;

import java.util.*;

import static org.junit.jupiter.api.Assertions.*;

public class RemapAppTest {
    @Test
    public void test1() {
        assertEquals(Map.of(1L, 2, 3L, 2), new RemapApp().remap(Map.of(2, List.of(1L, 3L))));
    }
}