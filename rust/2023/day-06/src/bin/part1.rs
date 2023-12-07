fn main() {
    let mut paper = Paper {
        time: Vec::with_capacity(4),
        distance: Vec::with_capacity(4),
    };
    include_str!("./input.txt")
        .lines()
        .enumerate()
        .for_each(|(i, line)| {
            let (_, nums) = line.split_once(":").unwrap();
            nums.split_whitespace()
                .map(|num| num.parse::<usize>().unwrap())
                .for_each(|parsed| match i {
                    0 => paper.time.push(parsed),
                    1 => paper.distance.push(parsed),
                    _ => panic!("unknown line {i}"),
                })
        });

    let product: usize = (0..paper.time.len())
        .map(|idx| (paper.time[idx], paper.distance[idx]))
        .map(|(time, distance)| count_possibilities(time, distance))
        .product();

    println!("{:?}", product);
}

struct Paper {
    time: Vec<usize>,
    distance: Vec<usize>,
}

fn count_possibilities(time: usize, distance: usize) -> usize {
    (0..=time)
        .filter_map(|ms| {
            let time_traveling = time - ms;
            let traveled_distace = time_traveling * ms;
            if traveled_distace <= distance {
                None
            } else {
                Some(traveled_distace)
            }
        })
        .count()
}
