# Game Introduction 

The Heros Idle Role Playing Game (RPG) is just what it sounds like: an RPG in which the players idle. In addition to just gaining levels, players can find items, go on quests, battle other players and find themselves in the games of Gods (some good and some evil). 

To play, you are only required to assign your Hero a name, a class and a race and register your Hero with the game server. The game server will monitor your Hero's adventures through the realm and update you of his endeavors. 


## Registering
To register, simply:

/msg IdleRPG REGISTER <char name> <password> <char class>

Where 'char name' can be up to 32 chars long, and 'char class' can be up to 32 chars.
Password can be any length, but keep it reasonable so the IRC server doesn't truncate it.

## Logging In
To login, simply:

/msg IdleRPG LOGIN <char name> <password>

This is a p0 (see Penalties) command.

## Logging Out
To logout, simply:


## Changing Your Class
To change your class, simply:

/msg IdleRPG NEWCLASS <new class>

This is a p0 (see Penalties) command.

## Removing Your Account
To remove your account, simply:

/msg IdleRPG REMOVEME

This is a p0 (see Penalties) command. :^)

## Changing Your Alignment
To change your alignment, simply:

/msg IdleRPG ALIGN <good|neutral|evil>

This is a p0 (see Penalties) command.

Your alignment can affect certain aspects of the game. You may align with good, neutral, or evil. 'Good' users have a 10% boost to their item sum for battles, and a 1/12 chance each day that they, along with a 'good' friend, will have the light of their god shine upon them, accelerating them 5-12% toward their next level. 'Evil' users have a 10% detriment to their item sum for battles (ever forsaken in their time of most need...), but have a 1/8 chance each day that they will either a) attempt to steal an item from a 'good' user (whom they cannot help but hate) or b) be forsaken (for 1-5% of their TTL) by their evil god. After all, we all know that crime doesn't pay. Also, 'good' users have only a 1/50 chance of landing a Critical Strike when battling, while 'evil' users (who always fight dirty) have a 1/20 chance. Neutral users haven't had anything changed, and all users start off as neutral.

I haven't run the numbers to see which alignment it is better to follow, so the stats for this feature may change in the future.

## Obtaining Bot Info
To see some simple information on the bot, simply:

/msg IdleRPG INFO

This is a p0 (see Penalties) command.

This command gives info such as to which server the bot is connected and the nicknames of online bot admins.

This command is optional, and may be disabled by your bot admin.

## Leveling
To gain levels, you must only be logged in and idle. The time between levels is based on your character level, and is calculated by the formula:

600 * (1.16 ^ YOUR_LEVEL)

Where ^ represents the exponentiation operator.

Very high levels are calculated differently. Levels after level 60 have a next time to level of:

(time to level @ 60) + ((1 day) * (level - 60))

The exponent method code had simply gotten to that point that levels were taking too long to complete.

Checking the Active Quest
To see the active quest, its users, and its time left to completion:

/msg IdleRPG QUEST

This is a p0 (see Penalties) command.

Checking Your Online Status
To see whether you are logged on, simply:

/msg IdleRPG WHOAMI

This is a p0 (see Penalties) command.

Penalties
If you do something other than idle, like part, quit, talk in the channel, change your nick, or notice the channel, you are penalized. The penalties are time, in seconds, added to your next time to level and are based on your character level. The formulae are as follows:

Nick change	30 * (1.14 ^ YOUR_LEVEL)
Part	200 * (1.14 ^ YOUR_LEVEL)
Quit	20 * (1.14 ^ YOUR_LEVEL)
LOGOUT command	20 * (1.14 ^ YOUR_LEVEL)
Being Kicked	250 * (1.14 ^ YOUR_LEVEL)
Channel privmsg	[message length] * (1.14 ^ YOUR_LEVEL)
Channel notice	[message length] * (1.14 ^ YOUR_LEVEL)
So, a level 25 character changing their nick would be penalized 20 * (1.14 ^ 25) = 793 seconds towards their next level.

Penalty shorthand is p[num]. So, a nick change is a p30 event, parting the channel is a p200 event, and quitting IRC is a p20 event. Messages and notices are p[length of message in characters].

Items
Each time you level, you find an item. You can find an item as high as 1.5 * YOUR_LEVEL (unless you find a unique item). There are 10 types of items: rings, amulets, charms, weapons, helms, tunics, gloves, leggings, shields, and boots. You can find one of each type. When you find an item with a level higher than the level of the item you already have, you toss the old item and start using the new one.

There is a p0 STATUS command which will show your item sum, but you cannot see which items you have over IRC. You can, however, see which items you have on the webhere.

As you may guess, you have a higher chance of rolling an item of a lower value than you do of rolling one of a higher value level. The exact formula is as follows:

for each 'number' from 1 to YOUR_LEVEL * 1.5
  you have a 1 / (1.4 ^ number) chance to find an item at this level
end for

As for item type, you have an equal chance to roll any type.

Battle
Each time you level, if your level is less than 25, you have a 25% chance to challenge someone to combat. If your level is greater than or equal to 25, you have a 100% chance to challenge someone. A pool of opponents is chosen of all online players, and one is chosen randomly. If there are no other online players, you fight no one. However, if you do challenge someone, this is how the victor is decided:

Your item levels are summed.
Their item levels are summed.
A random number between zero and your sum is taken.
A random number between zero and their sum is taken.
If your roll is larger than theirs, you win.
If you win, your time towards your next level is lowered. The amount that it is lowered is based on your opponent's level. The formula is:

((the larger number of (OPPONENT_LEVEL / 4) and 7) / 100) * YOUR_NEXT_TIME_TO_LEVEL

This means that you lose no less than 7% from your next time to level. If you win, your opponent is not penalized any time, unless you land a Critical Strike.

If you lose, you will be penalized time. The penalty is calculated using the formula:

((the larger number of (OPPONENT_LEVEL / 7) and 7) / 100) * YOUR_NEXT_TIME_TO_LEVEL

This means that you gain no less than 7% of your next time to level. If you lose, your opponent is not awarded any time.

Battling the IdleRPG bot is a special case. The bot has an item sum of 1+[highest item sum of all players]. The percent awarded if you win is a constant 20%, and the percent penalized if you lose is a constant 10%.

A successful battle may result an item being stolen.

Unique Items
After level 25, you have a chance to roll items significantly higher than items you would normally find at that level. These are unique items, and have the following stats:

Name	Item Level Range	Required User Level
Starlyte's Ring of Gyges	50-74	25 or greater
Shannon's Warm Wooly Gloves	50-74	25 or greater
Mike's Gore Amulet	75-99	30 or greater
Rahky's Obnoxious Charm of Contradiction	100-124	35 or greater
Murarth's 100% Evade Mantle	150-174	40 or greater
Hildy's Butterfly Sword	175-200	45 or greater
Bert's Blue Shoes	250-300	48 or greater
Rahx's Supersonic Leggings	300-350	50 or greater
John's Minion Shield	350-400	52 or greater
The chance of finding a unique item is:

(5 + YOUR_LEVEL * 0.25) in 100

The Hand of God
Every online user has roughly a 1/20 chance per day of a "Hand of God" affecting them. A HoG can help or hurt your character by carrying it between 5 and 75 percent towards or away from its next time to level. The odds are in your favor, however, with an 80% chance to help your character, and only a 20% chance of your character being smitten.

In addition to occurring randomly, admins may summon the HoG at their whim.

Critical Strike
If a challenger beats their opponent in battle, he has a chance of landing a Critical Strike. Neutral alignment players have a 1 / 35 chance, good alignment players have a 1 /50 chance, and evil alignment players have a 1 / 20 chance. If this occurs, his opponent is penalized time towards their next time to level.

This time will be no less than 5% and no more than 25% of their next time to level.

Team Battles
Every online user has roughly a 1/4 chance per day of being involved in a team battle. Team battles pit three online players against three other online players. Each side's items are summed, and a winner is chosen as in regular battling. If the first group bests the second group in combat, 20% of the lowest of the three's TTL is removed from their clocks. If the first group loses, 20% of their lowest member's TTL is added to their TTL.

## Calamities
Every online user has roughly a 1/8 chance per day of a calamity occurring to them. A calamity is a bit of extremely bad luck that either:

slows a player 5-12% of their next time to level
lowers one of their item's value by 10%

## Godsends
Every online user has roughly a 1/8 chance per day of a godsend occurring to them. A godsend is a bit of extremely good luck that either:

accelerates a player 5-12% of their next time to level
increases one of their item's value by 10%
Quests
Four randomly chosen players are chosen to represent and assist the Realm by going on a quest. If all four players make it to the quest's end, all questers are awarded by removing 25% of their TTL (ie, their TTL at quest's end). To complete a quest, no player can be penalized until the quest's end. If the quest is not completed, ALL online users are penalized a p15 as punishment.

## Item Stealing
After each battle, if the challenger wins, they have a 1 / 25 to steal an item from the challengee. Only items of a higher value are stolen, and the challenger's old item is given to the challengee in a moment of pity.

## Credits
Infinite thanks to the original creators of IdleRPG.