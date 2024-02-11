use std::fs;

fn main() {
    let content = fs::read_to_string("./day3-rucksack/input.txt").unwrap();

    let score: u16 = content.as_str()
        .split("\n")
        .map(|x| x.split_at(x.len() / 2))
        .map(|(a, b)| b
            .chars()
            .filter(|c| a.contains(c.to_owned()))
            .map(|c| {
                if c <= 'Z' {
                    (c as u16) - (b'A' as u16) + 27
                } else {
                    (c as u16) - (b'a' as u16) + 1
                }
            })
            .next()
            .unwrap()
        )
        .sum::<u16>();

    println!("{}", score);


    let score: u16 = content.as_str()
        .split("\n")
        .collect::<Vec<_>>()
        .chunks(3)
        .map(|set| set[0]
            .chars()
            .filter(|c| set[1].contains(c.to_owned()) && set[2].contains(c.to_owned()))
            .map(|c| {
                if c <= 'Z' {
                    (c as u16) - (b'A' as u16) + 27
                } else {
                    (c as u16) - (b'a' as u16) + 1
                }
            })
            .next()
            .unwrap()
        )
        .sum::<u16>();

    println!("{}", score);
}
