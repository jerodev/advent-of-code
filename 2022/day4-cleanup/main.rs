use std::fs;

fn main() {
    let content = fs::read_to_string("./day4-cleanup/input.txt").unwrap();

    let count = content.as_str()
        .split('\n')
        .filter(|s| {
            let (l, r) = s.split_once(',').unwrap();
            let pairs = (
                l.split_once('-').unwrap(),
                r.split_once('-').unwrap(),
            );

            let a = pairs.0.0.parse::<u8>().unwrap();
            let b = pairs.0.1.parse::<u8>().unwrap();
            let c = pairs.1.0.parse::<u8>().unwrap();
            let d = pairs.1.1.parse::<u8>().unwrap();

            (a >= c && b <= d) || (c >= a && d <= b)
        })
        .count();

    println!("{}", count);

    let count = content.as_str()
        .split('\n')
        .filter(|s| {
            let (l, r) = s.split_once(',').unwrap();
            let pairs = (
                l.split_once('-').unwrap(),
                r.split_once('-').unwrap(),
            );

            let a = pairs.0.0.parse::<u8>().unwrap();
            let b = pairs.0.1.parse::<u8>().unwrap();
            let c = pairs.1.0.parse::<u8>().unwrap();
            let d = pairs.1.1.parse::<u8>().unwrap();

            (a >= c && a <= d) || (b >= c && b <= d) || (a < c && b > d)
        })
        .count();

    println!("{}", count);
}
