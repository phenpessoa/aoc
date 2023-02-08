use std::error::Error;

const EXAMPLE_INPUT: &str = include_str!("example.txt");
const INPUT: &str = include_str!("input.txt");

pub fn do_magic(from: &str) -> Result<[isize; 2], Box<dyn Error>> {
    let data = match from {
        "example" => EXAMPLE_INPUT,
        "input" => INPUT,
        _ => return Err("unknown input".into()),
    };

    let mut vals = Vec::new();
    let mut cur = 0;

    for line in data.lines() {
        match line.parse::<isize>() {
            Ok(val) => cur += val,
            Err(_) => {
                vals.push(cur);
                cur = 0;
                continue;
            }
        }
    }

    if cur != 0 {
        vals.push(cur)
    }

    vals.sort();

    Ok([
        vals[vals.len() - 1],
        vals[vals.len() - 1] + vals[vals.len() - 2] + vals[vals.len() - 3],
    ])
}
