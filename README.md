# REAL TIME CHESS

Real time chess is terminal-based chess game where you've got to think fast. It
is intended to be played over a local network.

Opponnents do not take turns. Instead, they are limited only by the speed with
which they execute their desired moves.

You can download a ready to play version from the Releases section. Or you can
clone this repository.

## Setup
Real time chess runs in your terminal. There are two players, A and B. To start
a game, you need your and your opponents IP address. On linux systems, you can
find this with the following one liner `ip addr | grep 'state UP' -A2 | tail -n1 | awk '{print $2}' | cut -f1  -d'/'`. 
On Macs, try ` ipconfig getifaddr en0 ` (`en1` if using a wired connection.)

Now, as player A, let's say your ip address is `192.168.1.204` and your
opponent's is `192.168.1.208`. Issue the following commands:
```
export RTC_ID=A
export RTC_HostA='192.168.1.204'
export RTC_HostB='192.168.1.208'
```
Your opponent should do the same, but switch their RTC_ID to B. Now you're
ready to play!

## Gameplay
You can use the arrow or WASD keys to move your cursor. 'Return' switches from
select mode to insert mode. In insert mode, you can move your targeting
reticle and press 'Return' to move, or press 'Space' to return to select mode.

Pieces move according to chess rules, but you do not take turns. Develop your
pieces, go on the offensive, and DEATH TO THE KING!@!


## Contributing
Contributions extremely welcome! These can come in the form of filing and
issue, contributing a patch with fixes or new features, or starting
a discussion about something you'd like to see in Real Time Chess!

### Coming soon:
* Area damage and hitpoints
* Scorekeeping and game playback
* Automatic local network discovery/setup
* Chess Engine/Game AI
* Global network play
