fn main() {
    let _guard = sentry::init("https://7c9cdfb30fe04f248197a6f684d279ce@glitchtip-stg.puzzle.ch/9");
    panic!("Oh no, an error!");
}