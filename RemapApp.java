import java.util.*;

import static java.util.stream.Collectors.toMap;

public class RemapApp {
    public static void main(String[] args) {
        var data = new HashMap<Integer, List<Long>>() {{
            put(1, List.of(1L, 2L, 3L));
            put(2, List.of(11L, 22L));
        }};
        var remapped = data.entrySet().stream().flatMap(e -> e.getValue().stream().map(ee -> Map.entry(ee, e.getKey())))
            .collect(toMap(Map.Entry::getKey, Map.Entry::getValue));
    }
}
