fn main() {
    let sum: usize = include_str!("./input.txt")
        .lines()
        .map(|line| {
            let s = std::str::from_utf8(
                &line.as_bytes()[(line.find(":").unwrap()) + 1..],
            )
            .unwrap();

            s.split(";").map(|set| set.split(",")).map(|balls| {
                balls.map(|ball| {
                    let mut splitted_ball = ball.trim().split(" ");
                    let num =
                        splitted_ball.next().unwrap().parse::<usize>().unwrap();
                    match splitted_ball.next().unwrap() {
                        "red" => Ball::Red(num),
                        "green" => Ball::Green(num),
                        "blue" => Ball::Blue(num),
                        c => panic!("unknown color: {c}"),
                    }
                })
            })
        })
        .map(|game| {
            use std::cmp::max;
            let (min_red, min_green, min_blue) = game.flatten().fold(
                (1, 1, 1),
                |(min_red, min_green, min_blue), ball| match ball {
                    Ball::Red(n) => (max(min_red, n), min_green, min_blue),
                    Ball::Green(n) => (min_red, max(min_green, n), min_blue),
                    Ball::Blue(n) => (min_red, min_green, max(min_blue, n)),
                },
            );
            min_red * min_green * min_blue
        })
        .sum();
    println!("{sum}")
}

enum Ball {
    Red(usize),
    Green(usize),
    Blue(usize),
}
