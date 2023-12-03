fn main() {
    let sum: usize = include_str!("./input.txt")
        .lines()
        .enumerate()
        .map(|(i, line)| {
            let s = std::str::from_utf8(
                &line.as_bytes()[(line.find(":").unwrap()) + 1..],
            )
            .unwrap();
            (
                i,
                s.split(";").map(|set| set.split(",")).map(|balls| {
                    balls.map(|ball| {
                        let mut splitted_ball = ball.trim().split(" ");
                        let num = splitted_ball
                            .next()
                            .unwrap()
                            .parse::<usize>()
                            .unwrap();
                        match splitted_ball.next().unwrap() {
                            "red" => Ball::Red(num),
                            "green" => Ball::Green(num),
                            "blue" => Ball::Blue(num),
                            c => panic!("unknown color: {c}"),
                        }
                    })
                }),
            )
        })
        .filter_map(|(i, game)| {
            game.flatten()
                .map(|ball| match ball {
                    Ball::Red(n) if n > 12 => false,
                    Ball::Green(n) if n > 13 => false,
                    Ball::Blue(n) if n > 14 => false,
                    _ => true,
                })
                .all(|_bool| _bool)
                .then_some(i + 1)
        })
        .sum();
    println!("{sum}")
}

enum Ball {
    Red(usize),
    Green(usize),
    Blue(usize),
}
