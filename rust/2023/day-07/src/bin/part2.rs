use std::collections::HashMap;

fn main() {
    let mut hands: Vec<_> = include_str!("./input.txt")
        .lines()
        .map(|line| line.split_once(" ").unwrap())
        .map(|(hand, bid)| parse_hand(hand, bid))
        .collect();
    hands.sort();
    let total_winnings: usize = hands
        .iter()
        .enumerate()
        .map(|(i, hand)| hand.bid * (i + 1))
        .sum();
    println!("{total_winnings}");
}

fn parse_hand(hand_str: &str, bid_str: &str) -> Hand {
    let mut map: HashMap<char, usize> = HashMap::new();
    hand_str
        .chars()
        .for_each(|c| *map.entry(c).or_insert(0) += 1);
    let mut counts: Vec<_> = map.iter().map(|(_, count)| *count).collect();
    counts.sort();
    counts.reverse();

    let bytes = hand_str.as_bytes();
    let cards: [Card; 5] = [
        Card::new(bytes[0]),
        Card::new(bytes[1]),
        Card::new(bytes[2]),
        Card::new(bytes[3]),
        Card::new(bytes[4]),
    ];

    let jokers_count = cards.iter().filter(|&card| *card == Card::Jack).count();

    let game = match counts[0] {
        5 => Game::FiveOfAKind,
        4 => Game::FourOfAKind,
        3 if counts[1] == 2 => Game::FullHouse,
        3 => Game::ThreeOfAKind,
        2 if counts[1] == 2 => Game::TwoPairs,
        2 => Game::Pair,
        _ => Game::HighCard,
    };

    let bid = bid_str.trim().parse::<usize>().unwrap();

    if jokers_count == 0 {
        return Hand { game, cards, bid };
    }

    let game = match (game, jokers_count) {
        (Game::FiveOfAKind, _) => Game::FiveOfAKind,

        (Game::FourOfAKind, 1) => Game::FiveOfAKind,
        (Game::FourOfAKind, 4) => Game::FiveOfAKind,

        (Game::FullHouse, 2) => Game::FiveOfAKind,
        (Game::FullHouse, 3) => Game::FiveOfAKind,

        (Game::ThreeOfAKind, 2) => Game::FiveOfAKind,
        (Game::ThreeOfAKind, 3) => Game::FourOfAKind,
        (Game::ThreeOfAKind, 1) => Game::FourOfAKind,

        (Game::TwoPairs, 2) => Game::FourOfAKind,
        (Game::TwoPairs, 1) => Game::FullHouse,

        (Game::Pair, 2) => Game::ThreeOfAKind,
        (Game::Pair, 1) => Game::ThreeOfAKind,

        (Game::HighCard, 1) => Game::Pair,

        (g, _) => g,
    };

    Hand { game, cards, bid }
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord)]
enum Card {
    Jack,
    Two,
    Three,
    Four,
    Five,
    Six,
    Seven,
    Eight,
    Nine,
    Ten,
    Queen,
    King,
    Ace,
}

impl Card {
    fn new(c: u8) -> Card {
        match c {
            b'2' => Card::Two,
            b'3' => Card::Three,
            b'4' => Card::Four,
            b'5' => Card::Five,
            b'6' => Card::Six,
            b'7' => Card::Seven,
            b'8' => Card::Eight,
            b'9' => Card::Nine,
            b'T' => Card::Ten,
            b'J' => Card::Jack,
            b'Q' => Card::Queen,
            b'K' => Card::King,
            b'A' => Card::Ace,
            _ => panic!("unknown card"),
        }
    }
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord)]
struct Hand {
    game: Game,
    cards: [Card; 5],
    bid: usize,
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord)]
enum Game {
    HighCard,
    Pair,
    TwoPairs,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind,
}
