fn main() {
    let input = include_str!("./input.txt");
    let mut time_strs: Vec<_> = Vec::with_capacity(4);
    let mut distance_strs: Vec<_> = Vec::with_capacity(4);

    input.lines().enumerate().for_each(|(i, line)| {
        let (_, nums) = line.split_once(":").unwrap();
        nums.split_whitespace().for_each(|num| match i {
            0 => time_strs.push(num),
            1 => distance_strs.push(num),
            _ => panic!("unknown line {i}"),
        })
    });

    let paper = Paper {
        time: time_strs.join("").parse::<usize>().unwrap(),
        distance: distance_strs.join("").parse::<usize>().unwrap(),
    };

    println!("{}", count_possibilities(paper.time, paper.distance));
}

struct Paper {
    time: usize,
    distance: usize,
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
