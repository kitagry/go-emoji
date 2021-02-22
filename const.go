package emoji

type Emoji struct {
	emoji  []byte
	markup []byte
}

var emojis []Emoji = []Emoji{
	{[]byte("😊"), []byte(":blush:")},
	{[]byte("😃"), []byte(":smiley:")},
	{[]byte("☺️"), []byte(":relaxed:")},
	{[]byte("😏"), []byte(":smirk:")},
	{[]byte("😍"), []byte(":heart_eyes:")},
	{[]byte("😘"), []byte(":kissing_heart:")},
	{[]byte("😚"), []byte(":kissing_closed_eyes:")},
	{[]byte("😳"), []byte(":flushed:")},
	{[]byte("😌"), []byte(":relieved:")},
	{[]byte("😆"), []byte(":satisfied:")},
	{[]byte("😁"), []byte(":grin:")},
	{[]byte("😉"), []byte(":wink:")},
	{[]byte("😜"), []byte(":stuck_out_tongue_winking_eye:")},
	{[]byte("😝"), []byte(":stuck_out_tongue_closed_eyes:")},
	{[]byte("😀"), []byte(":grinning:")},
	{[]byte("😗"), []byte(":kissing:")},
	{[]byte("😙"), []byte(":kissing_smiling_eyes:")},
	{[]byte("😛"), []byte(":stuck_out_tongue:")},
	{[]byte("😴"), []byte(":sleeping:")},
	{[]byte("😟"), []byte(":worried:")},
	{[]byte("😦"), []byte(":frowning:")},
	{[]byte("😧"), []byte(":anguished:")},
	{[]byte("😮"), []byte(":open_mouth:")},
	{[]byte("😬"), []byte(":grimacing:")},
	{[]byte("😕"), []byte(":confused:")},
	{[]byte("😯"), []byte(":hushed:")},
	{[]byte("😑"), []byte(":expressionless:")},
	{[]byte("😒"), []byte(":unamused:")},
	{[]byte("😅"), []byte(":sweat_smile:")},
	{[]byte("😓"), []byte(":sweat:")},
	{[]byte("😥"), []byte(":disappointed_relieved:")},
	{[]byte("😩"), []byte(":weary:")},
	{[]byte("😔"), []byte(":pensive:")},
	{[]byte("😞"), []byte(":disappointed:")},
	{[]byte("😖"), []byte(":confounded:")},
	{[]byte("😨"), []byte(":fearful:")},
	{[]byte("😰"), []byte(":cold_sweat:")},
	{[]byte("😣"), []byte(":persevere:")},
	{[]byte("😢"), []byte(":cry:")},
	{[]byte("😭"), []byte(":sob:")},
	{[]byte("😂"), []byte(":joy:")},
	{[]byte("😲"), []byte(":astonished:")},
	{[]byte("😱"), []byte(":scream:")},
	{[]byte("😫"), []byte(":tired_face:")},
	{[]byte("😠"), []byte(":angry:")},
	{[]byte("😡"), []byte(":rage:")},
	{[]byte("😤"), []byte(":triumph:")},
	{[]byte("😪"), []byte(":sleepy:")},
	{[]byte("😋"), []byte(":yum:")},
	{[]byte("😷"), []byte(":mask:")},
	{[]byte("😎"), []byte(":sunglasses:")},
	{[]byte("😵"), []byte(":dizzy_face:")},
	{[]byte("👿"), []byte(":imp:")},
	{[]byte("😈"), []byte(":smiling_imp:")},
	{[]byte("😐"), []byte(":neutral_face:")},
	{[]byte("😶"), []byte(":no_mouth:")},
	{[]byte("😇"), []byte(":innocent:")},
	{[]byte("👽"), []byte(":alien:")},
	{[]byte("💛"), []byte(":yellow_heart:")},
	{[]byte("💙"), []byte(":blue_heart:")},
	{[]byte("💜"), []byte(":purple_heart:")},
	{[]byte("❤️"), []byte(":heart:")},
	{[]byte("💚"), []byte(":green_heart:")},
	{[]byte("💔"), []byte(":broken_heart:")},
	{[]byte("💓"), []byte(":heartbeat:")},
	{[]byte("💗"), []byte(":heartpulse:")},
	{[]byte("💕"), []byte(":two_hearts:")},
	{[]byte("💞"), []byte(":revolving_hearts:")},
	{[]byte("💘"), []byte(":cupid:")},
	{[]byte("💖"), []byte(":sparkling_heart:")},
	{[]byte("✨"), []byte(":sparkles:")},
	{[]byte("⭐"), []byte(":star:")},
	{[]byte("🌟"), []byte(":star2:")},
	{[]byte("💫"), []byte(":dizzy:")},
	{[]byte("💥"), []byte(":boom:")},
	{[]byte("💥"), []byte(":collision:")},
	{[]byte("💢"), []byte(":anger:")},
	{[]byte("❗"), []byte(":exclamation:")},
	{[]byte("❓"), []byte(":question:")},
	{[]byte("❕"), []byte(":grey_exclamation:")},
	{[]byte("❔"), []byte(":grey_question:")},
	{[]byte("💤"), []byte(":zzz:")},
	{[]byte("💨"), []byte(":dash:")},
	{[]byte("💦"), []byte(":sweat_drops:")},
	{[]byte("🎶"), []byte(":notes:")},
	{[]byte("🎵"), []byte(":musical_note:")},
	{[]byte("🔥"), []byte(":fire:")},
	{[]byte("💩"), []byte(":hankey:")},
	{[]byte("💩"), []byte(":poop:")},
	{[]byte("💩"), []byte(":shit:")},
	{[]byte("👍"), []byte(":+1:")},
	{[]byte("👍"), []byte(":thumbsup:")},
	{[]byte("👎"), []byte(":-1:")},
	{[]byte("👎"), []byte(":thumbsdown:")},
	{[]byte("👌"), []byte(":ok_hand:")},
	{[]byte("👊"), []byte(":punch:")},
	{[]byte("👊"), []byte(":facepunch:")},
	{[]byte("✊"), []byte(":fist:")},
	{[]byte("✌️"), []byte(":v:")},
	{[]byte("👋"), []byte(":wave:")},
	{[]byte("✋"), []byte(":hand:")},
	{[]byte("✋"), []byte(":raised_hand:")},
	{[]byte("👐"), []byte(":open_hands:")},
	{[]byte("☝️"), []byte(":point_up:")},
	{[]byte("👇"), []byte(":point_down:")},
	{[]byte("👈"), []byte(":point_left:")},
	{[]byte("👉"), []byte(":point_right:")},
	{[]byte("🙌"), []byte(":raised_hands:")},
	{[]byte("🙏"), []byte(":pray:")},
	{[]byte("👆"), []byte(":point_up_2:")},
	{[]byte("👏"), []byte(":clap:")},
	{[]byte("💪"), []byte(":muscle:")},
	{[]byte("🤘"), []byte(":metal:")},
	{[]byte("🖕"), []byte(":fu:")},
	{[]byte("🚶"), []byte(":walking:")},
	{[]byte("🏃"), []byte(":runner:")},
	{[]byte("🏃"), []byte(":running:")},
	{[]byte("👫"), []byte(":couple:")},
	{[]byte("👪"), []byte(":family:")},
	{[]byte("👬"), []byte(":two_men_holding_hands:")},
	{[]byte("👭"), []byte(":two_women_holding_hands:")},
	{[]byte("💃"), []byte(":dancer:")},
	{[]byte("👯"), []byte(":dancers:")},
	{[]byte("🙆‍♀️"), []byte(":ok_woman:")},
	{[]byte("🙅"), []byte(":no_good:")},
	{[]byte("💁"), []byte(":information_desk_person:")},
	{[]byte("🙋"), []byte(":raising_hand:")},
	{[]byte("👰"), []byte(":bride_with_veil:")},
	{[]byte("🙇"), []byte(":bow:")},
	{[]byte("💏"), []byte(":couplekiss:")},
	{[]byte("💑"), []byte(":couple_with_heart:")},
	{[]byte("💆"), []byte(":massage:")},
	{[]byte("💇"), []byte(":haircut:")},
	{[]byte("💅"), []byte(":nail_care:")},
	{[]byte("👦"), []byte(":boy:")},
	{[]byte("👧"), []byte(":girl:")},
	{[]byte("👩"), []byte(":woman:")},
	{[]byte("👨"), []byte(":man:")},
	{[]byte("👶"), []byte(":baby:")},
	{[]byte("👵"), []byte(":older_woman:")},
	{[]byte("👴"), []byte(":older_man:")},
	{[]byte("👲"), []byte(":man_with_gua_pi_mao:")},
	{[]byte("👳‍♂️"), []byte(":man_with_turban:")},
	{[]byte("👷"), []byte(":construction_worker:")},
	{[]byte("👮"), []byte(":cop:")},
	{[]byte("👼"), []byte(":angel:")},
	{[]byte("👸"), []byte(":princess:")},
	{[]byte("😺"), []byte(":smiley_cat:")},
	{[]byte("😸"), []byte(":smile_cat:")},
	{[]byte("😻"), []byte(":heart_eyes_cat:")},
	{[]byte("😽"), []byte(":kissing_cat:")},
	{[]byte("😼"), []byte(":smirk_cat:")},
	{[]byte("🙀"), []byte(":scream_cat:")},
	{[]byte("😿"), []byte(":crying_cat_face:")},
	{[]byte("😹"), []byte(":joy_cat:")},
	{[]byte("😾"), []byte(":pouting_cat:")},
	{[]byte("👹"), []byte(":japanese_ogre:")},
	{[]byte("👺"), []byte(":japanese_goblin:")},
	{[]byte("🙈"), []byte(":see_no_evil:")},
	{[]byte("🙉"), []byte(":hear_no_evil:")},
	{[]byte("🙊"), []byte(":speak_no_evil:")},
	{[]byte("💂‍♂️"), []byte(":guardsman:")},
	{[]byte("💀"), []byte(":skull:")},
	{[]byte("🐾"), []byte(":feet:")},
	{[]byte("👄"), []byte(":lips:")},
	{[]byte("💋"), []byte(":kiss:")},
	{[]byte("💧"), []byte(":droplet:")},
	{[]byte("👂"), []byte(":ear:")},
	{[]byte("👀"), []byte(":eyes:")},
	{[]byte("👃"), []byte(":nose:")},
	{[]byte("👅"), []byte(":tongue:")},
	{[]byte("💌"), []byte(":love_letter:")},
	{[]byte("👤"), []byte(":bust_in_silhouette:")},
	{[]byte("👥"), []byte(":busts_in_silhouette:")},
	{[]byte("💬"), []byte(":speech_balloon:")},
	{[]byte("💭"), []byte(":thought_balloon:")},
	{[]byte("❄️"), []byte(":snowflake:")},
	{[]byte("⛄"), []byte(":snowman:")},
	{[]byte("⚡"), []byte(":zap:")},
	{[]byte("🌀"), []byte(":cyclone:")},
	{[]byte("🌁"), []byte(":foggy:")},
	{[]byte("🌊"), []byte(":ocean:")},
	{[]byte("🐱"), []byte(":cat:")},
	{[]byte("🐶"), []byte(":dog:")},
	{[]byte("🐭"), []byte(":mouse:")},
	{[]byte("🐹"), []byte(":hamster:")},
	{[]byte("🐰"), []byte(":rabbit:")},
	{[]byte("🐺"), []byte(":wolf:")},
	{[]byte("🐸"), []byte(":frog:")},
	{[]byte("🐯"), []byte(":tiger:")},
	{[]byte("🐨"), []byte(":koala:")},
	{[]byte("🐻"), []byte(":bear:")},
	{[]byte("🐷"), []byte(":pig:")},
	{[]byte("🐽"), []byte(":pig_nose:")},
	{[]byte("🐮"), []byte(":cow:")},
	{[]byte("🐗"), []byte(":boar:")},
	{[]byte("🐵"), []byte(":monkey_face:")},
	{[]byte("🐒"), []byte(":monkey:")},
	{[]byte("🐴"), []byte(":horse:")},
	{[]byte("🐎"), []byte(":racehorse:")},
	{[]byte("🐫"), []byte(":camel:")},
	{[]byte("🐑"), []byte(":sheep:")},
	{[]byte("🐘"), []byte(":elephant:")},
	{[]byte("🐼"), []byte(":panda_face:")},
	{[]byte("🐍"), []byte(":snake:")},
	{[]byte("🐦"), []byte(":bird:")},
	{[]byte("🐤"), []byte(":baby_chick:")},
	{[]byte("🐥"), []byte(":hatched_chick:")},
	{[]byte("🐣"), []byte(":hatching_chick:")},
	{[]byte("🐔"), []byte(":chicken:")},
	{[]byte("🐧"), []byte(":penguin:")},
	{[]byte("🐢"), []byte(":turtle:")},
	{[]byte("🐛"), []byte(":bug:")},
	{[]byte("🐝"), []byte(":honeybee:")},
	{[]byte("🐜"), []byte(":ant:")},
	{[]byte("🐞"), []byte(":beetle:")},
	{[]byte("🐌"), []byte(":snail:")},
	{[]byte("🐙"), []byte(":octopus:")},
	{[]byte("🐠"), []byte(":tropical_fish:")},
	{[]byte("🐟"), []byte(":fish:")},
	{[]byte("🐳"), []byte(":whale:")},
	{[]byte("🐋"), []byte(":whale2:")},
	{[]byte("🐬"), []byte(":dolphin:")},
	{[]byte("🐄"), []byte(":cow2:")},
	{[]byte("🐏"), []byte(":ram:")},
	{[]byte("🐀"), []byte(":rat:")},
	{[]byte("🐃"), []byte(":water_buffalo:")},
	{[]byte("🐅"), []byte(":tiger2:")},
	{[]byte("🐇"), []byte(":rabbit2:")},
	{[]byte("🐉"), []byte(":dragon:")},
	{[]byte("🐐"), []byte(":goat:")},
	{[]byte("🐓"), []byte(":rooster:")},
	{[]byte("🐕"), []byte(":dog2:")},
	{[]byte("🐖"), []byte(":pig2:")},
	{[]byte("🐁"), []byte(":mouse2:")},
	{[]byte("🐂"), []byte(":ox:")},
	{[]byte("🐲"), []byte(":dragon_face:")},
	{[]byte("🐡"), []byte(":blowfish:")},
	{[]byte("🐊"), []byte(":crocodile:")},
	{[]byte("🐪"), []byte(":dromedary_camel:")},
	{[]byte("🐆"), []byte(":leopard:")},
	{[]byte("🐈"), []byte(":cat2:")},
	{[]byte("🐩"), []byte(":poodle:")},
	{[]byte("🐾"), []byte(":paw_prints:")},
	{[]byte("💐"), []byte(":bouquet:")},
	{[]byte("🌸"), []byte(":cherry_blossom:")},
	{[]byte("🌷"), []byte(":tulip:")},
	{[]byte("🍀"), []byte(":four_leaf_clover:")},
	{[]byte("🌹"), []byte(":rose:")},
	{[]byte("🌻"), []byte(":sunflower:")},
	{[]byte("🌺"), []byte(":hibiscus:")},
	{[]byte("🍁"), []byte(":maple_leaf:")},
	{[]byte("🍃"), []byte(":leaves:")},
	{[]byte("🍂"), []byte(":fallen_leaf:")},
	{[]byte("🌿"), []byte(":herb:")},
	{[]byte("🍄"), []byte(":mushroom:")},
	{[]byte("🌵"), []byte(":cactus:")},
	{[]byte("🌴"), []byte(":palm_tree:")},
	{[]byte("🌲"), []byte(":evergreen_tree:")},
	{[]byte("🌳"), []byte(":deciduous_tree:")},
	{[]byte("🌰"), []byte(":chestnut:")},
	{[]byte("🌱"), []byte(":seedling:")},
	{[]byte("🌼"), []byte(":blossom:")},
	{[]byte("🌾"), []byte(":ear_of_rice:")},
	{[]byte("🐚"), []byte(":shell:")},
	{[]byte("🌐"), []byte(":globe_with_meridians:")},
	{[]byte("🌞"), []byte(":sun_with_face:")},
	{[]byte("🌝"), []byte(":full_moon_with_face:")},
	{[]byte("🌚"), []byte(":new_moon_with_face:")},
	{[]byte("🌑"), []byte(":new_moon:")},
	{[]byte("🌒"), []byte(":waxing_crescent_moon:")},
	{[]byte("🌓"), []byte(":first_quarter_moon:")},
	{[]byte("🌔"), []byte(":waxing_gibbous_moon:")},
	{[]byte("🌕"), []byte(":full_moon:")},
	{[]byte("🌖"), []byte(":waning_gibbous_moon:")},
	{[]byte("🌗"), []byte(":last_quarter_moon:")},
	{[]byte("🌘"), []byte(":waning_crescent_moon:")},
	{[]byte("🌜"), []byte(":last_quarter_moon_with_face:")},
	{[]byte("🌛"), []byte(":first_quarter_moon_with_face:")},
	{[]byte("🌔"), []byte(":moon:")},
	{[]byte("🌍"), []byte(":earth_africa:")},
	{[]byte("🌎"), []byte(":earth_americas:")},
	{[]byte("🌏"), []byte(":earth_asia:")},
	{[]byte("🌋"), []byte(":volcano:")},
	{[]byte("🌌"), []byte(":milky_way:")},
	{[]byte("⛅"), []byte(":partly_sunny:")},
	{[]byte("🎒"), []byte(":school_satchel:")},
	{[]byte("🎓"), []byte(":mortar_board:")},
	{[]byte("🎏"), []byte(":flags:")},
	{[]byte("🎆"), []byte(":fireworks:")},
	{[]byte("🎇"), []byte(":sparkler:")},
	{[]byte("🎐"), []byte(":wind_chime:")},
	{[]byte("🎑"), []byte(":rice_scene:")},
	{[]byte("🎃"), []byte(":jack_o_lantern:")},
	{[]byte("👻"), []byte(":ghost:")},
	{[]byte("🎅"), []byte(":santa:")},
	{[]byte("🎄"), []byte(":christmas_tree:")},
	{[]byte("🎁"), []byte(":gift:")},
	{[]byte("🔔"), []byte(":bell:")},
	{[]byte("🔕"), []byte(":no_bell:")},
	{[]byte("🎋"), []byte(":tanabata_tree:")},
	{[]byte("🎉"), []byte(":tada:")},
	{[]byte("🎊"), []byte(":confetti_ball:")},
	{[]byte("🎈"), []byte(":balloon:")},
	{[]byte("🔮"), []byte(":crystal_ball:")},
	{[]byte("💿"), []byte(":cd:")},
	{[]byte("📀"), []byte(":dvd:")},
	{[]byte("💾"), []byte(":floppy_disk:")},
	{[]byte("📷"), []byte(":camera:")},
	{[]byte("📹"), []byte(":video_camera:")},
	{[]byte("🎥"), []byte(":movie_camera:")},
	{[]byte("💻"), []byte(":computer:")},
	{[]byte("📺"), []byte(":tv:")},
	{[]byte("📱"), []byte(":iphone:")},
	{[]byte("☎️"), []byte(":phone:")},
	{[]byte("☎️"), []byte(":telephone:")},
	{[]byte("📞"), []byte(":telephone_receiver:")},
	{[]byte("📟"), []byte(":pager:")},
	{[]byte("📠"), []byte(":fax:")},
	{[]byte("💽"), []byte(":minidisc:")},
	{[]byte("📼"), []byte(":vhs:")},
	{[]byte("🔉"), []byte(":sound:")},
	{[]byte("🔈"), []byte(":speaker:")},
	{[]byte("🔇"), []byte(":mute:")},
	{[]byte("📢"), []byte(":loudspeaker:")},
	{[]byte("📣"), []byte(":mega:")},
	{[]byte("⌛"), []byte(":hourglass:")},
	{[]byte("⏳"), []byte(":hourglass_flowing_sand:")},
	{[]byte("⏰"), []byte(":alarm_clock:")},
	{[]byte("⌚"), []byte(":watch:")},
	{[]byte("📻"), []byte(":radio:")},
	{[]byte("📡"), []byte(":satellite:")},
	{[]byte("➿"), []byte(":loop:")},
	{[]byte("🔍"), []byte(":mag:")},
	{[]byte("🔎"), []byte(":mag_right:")},
	{[]byte("🔓"), []byte(":unlock:")},
	{[]byte("🔒"), []byte(":lock:")},
	{[]byte("🔏"), []byte(":lock_with_ink_pen:")},
	{[]byte("🔐"), []byte(":closed_lock_with_key:")},
	{[]byte("🔑"), []byte(":key:")},
	{[]byte("💡"), []byte(":bulb:")},
	{[]byte("🔦"), []byte(":flashlight:")},
	{[]byte("🔆"), []byte(":high_brightness:")},
	{[]byte("🔅"), []byte(":low_brightness:")},
	{[]byte("🔌"), []byte(":electric_plug:")},
	{[]byte("🔋"), []byte(":battery:")},
	{[]byte("📲"), []byte(":calling:")},
	{[]byte("✉️"), []byte(":email:")},
	{[]byte("📫"), []byte(":mailbox:")},
	{[]byte("📮"), []byte(":postbox:")},
	{[]byte("🛀"), []byte(":bath:")},
	{[]byte("🛁"), []byte(":bathtub:")},
	{[]byte("🚿"), []byte(":shower:")},
	{[]byte("🚽"), []byte(":toilet:")},
	{[]byte("🔧"), []byte(":wrench:")},
	{[]byte("🔩"), []byte(":nut_and_bolt:")},
	{[]byte("🔨"), []byte(":hammer:")},
	{[]byte("💺"), []byte(":seat:")},
	{[]byte("💰"), []byte(":moneybag:")},
	{[]byte("💴"), []byte(":yen:")},
	{[]byte("💵"), []byte(":dollar:")},
	{[]byte("💷"), []byte(":pound:")},
	{[]byte("💶"), []byte(":euro:")},
	{[]byte("💳"), []byte(":credit_card:")},
	{[]byte("💸"), []byte(":money_with_wings:")},
	{[]byte("📧"), []byte(":e-mail:")},
	{[]byte("📥"), []byte(":inbox_tray:")},
	{[]byte("📤"), []byte(":outbox_tray:")},
	{[]byte("✉️"), []byte(":envelope:")},
	{[]byte("📨"), []byte(":incoming_envelope:")},
	{[]byte("📯"), []byte(":postal_horn:")},
	{[]byte("📪"), []byte(":mailbox_closed:")},
	{[]byte("📬"), []byte(":mailbox_with_mail:")},
	{[]byte("📭"), []byte(":mailbox_with_no_mail:")},
	{[]byte("🚪"), []byte(":door:")},
	{[]byte("🚬"), []byte(":smoking:")},
	{[]byte("💣"), []byte(":bomb:")},
	{[]byte("🔫"), []byte(":gun:")},
	{[]byte("🔪"), []byte(":hocho:")},
	{[]byte("💊"), []byte(":pill:")},
	{[]byte("💉"), []byte(":syringe:")},
	{[]byte("📄"), []byte(":page_facing_up:")},
	{[]byte("📃"), []byte(":page_with_curl:")},
	{[]byte("📑"), []byte(":bookmark_tabs:")},
	{[]byte("📊"), []byte(":bar_chart:")},
	{[]byte("📈"), []byte(":chart_with_upwards_trend:")},
	{[]byte("📉"), []byte(":chart_with_downwards_trend:")},
	{[]byte("📜"), []byte(":scroll:")},
	{[]byte("📋"), []byte(":clipboard:")},
	{[]byte("📆"), []byte(":calendar:")},
	{[]byte("📅"), []byte(":date:")},
	{[]byte("📇"), []byte(":card_index:")},
	{[]byte("📁"), []byte(":file_folder:")},
	{[]byte("📂"), []byte(":open_file_folder:")},
	{[]byte("✂️"), []byte(":scissors:")},
	{[]byte("📌"), []byte(":pushpin:")},
	{[]byte("📎"), []byte(":paperclip:")},
	{[]byte("✒️"), []byte(":black_nib:")},
	{[]byte("✏️"), []byte(":pencil2:")},
	{[]byte("📏"), []byte(":straight_ruler:")},
	{[]byte("📐"), []byte(":triangular_ruler:")},
	{[]byte("📕"), []byte(":closed_book:")},
	{[]byte("📗"), []byte(":green_book:")},
	{[]byte("📘"), []byte(":blue_book:")},
	{[]byte("📙"), []byte(":orange_book:")},
	{[]byte("📓"), []byte(":notebook:")},
	{[]byte("📔"), []byte(":notebook_with_decorative_cover:")},
	{[]byte("📒"), []byte(":ledger:")},
	{[]byte("📚"), []byte(":books:")},
	{[]byte("🔖"), []byte(":bookmark:")},
	{[]byte("📛"), []byte(":name_badge:")},
	{[]byte("🔬"), []byte(":microscope:")},
	{[]byte("🔭"), []byte(":telescope:")},
	{[]byte("📰"), []byte(":newspaper:")},
	{[]byte("🏈"), []byte(":football:")},
	{[]byte("🏀"), []byte(":basketball:")},
	{[]byte("⚽"), []byte(":soccer:")},
	{[]byte("⚾"), []byte(":baseball:")},
	{[]byte("🎾"), []byte(":tennis:")},
	{[]byte("🎱"), []byte(":8ball:")},
	{[]byte("🏉"), []byte(":rugby_football:")},
	{[]byte("🎳"), []byte(":bowling:")},
	{[]byte("⛳"), []byte(":golf:")},
	{[]byte("🚵"), []byte(":mountain_bicyclist:")},
	{[]byte("🚴"), []byte(":bicyclist:")},
	{[]byte("🏇"), []byte(":horse_racing:")},
	{[]byte("🏂"), []byte(":snowboarder:")},
	{[]byte("🏊"), []byte(":swimmer:")},
	{[]byte("🏄"), []byte(":surfer:")},
	{[]byte("🎿"), []byte(":ski:")},
	{[]byte("♠️"), []byte(":spades:")},
	{[]byte("♥️"), []byte(":hearts:")},
	{[]byte("♣️"), []byte(":clubs:")},
	{[]byte("♦️"), []byte(":diamonds:")},
	{[]byte("💎"), []byte(":gem:")},
	{[]byte("💍"), []byte(":ring:")},
	{[]byte("🏆"), []byte(":trophy:")},
	{[]byte("🎼"), []byte(":musical_score:")},
	{[]byte("🎹"), []byte(":musical_keyboard:")},
	{[]byte("🎻"), []byte(":violin:")},
	{[]byte("👾"), []byte(":space_invader:")},
	{[]byte("🎮"), []byte(":video_game:")},
	{[]byte("🃏"), []byte(":black_joker:")},
	{[]byte("🎴"), []byte(":flower_playing_cards:")},
	{[]byte("🎲"), []byte(":game_die:")},
	{[]byte("🎯"), []byte(":dart:")},
	{[]byte("🀄"), []byte(":mahjong:")},
	{[]byte("🎬"), []byte(":clapper:")},
	{[]byte("📝"), []byte(":memo:")},
	{[]byte("📝"), []byte(":pencil:")},
	{[]byte("📖"), []byte(":book:")},
	{[]byte("🎨"), []byte(":art:")},
	{[]byte("🎤"), []byte(":microphone:")},
	{[]byte("🎧"), []byte(":headphones:")},
	{[]byte("🎺"), []byte(":trumpet:")},
	{[]byte("🎷"), []byte(":saxophone:")},
	{[]byte("🎸"), []byte(":guitar:")},
	{[]byte("👞"), []byte(":shoe:")},
	{[]byte("👡"), []byte(":sandal:")},
	{[]byte("👠"), []byte(":high_heel:")},
	{[]byte("💄"), []byte(":lipstick:")},
	{[]byte("👢"), []byte(":boot:")},
	{[]byte("👕"), []byte(":shirt:")},
	{[]byte("👕"), []byte(":tshirt:")},
	{[]byte("👔"), []byte(":necktie:")},
	{[]byte("👚"), []byte(":womans_clothes:")},
	{[]byte("👗"), []byte(":dress:")},
	{[]byte("🎽"), []byte(":running_shirt_with_sash:")},
	{[]byte("👖"), []byte(":jeans:")},
	{[]byte("👘"), []byte(":kimono:")},
	{[]byte("👙"), []byte(":bikini:")},
	{[]byte("🎀"), []byte(":ribbon:")},
	{[]byte("🎩"), []byte(":tophat:")},
	{[]byte("👑"), []byte(":crown:")},
	{[]byte("👒"), []byte(":womans_hat:")},
	{[]byte("👞"), []byte(":mans_shoe:")},
	{[]byte("🌂"), []byte(":closed_umbrella:")},
	{[]byte("💼"), []byte(":briefcase:")},
	{[]byte("👜"), []byte(":handbag:")},
	{[]byte("👝"), []byte(":pouch:")},
	{[]byte("👛"), []byte(":purse:")},
	{[]byte("👓"), []byte(":eyeglasses:")},
	{[]byte("🎣"), []byte(":fishing_pole_and_fish:")},
	{[]byte("☕"), []byte(":coffee:")},
	{[]byte("🍵"), []byte(":tea:")},
	{[]byte("🍶"), []byte(":sake:")},
	{[]byte("🍼"), []byte(":baby_bottle:")},
	{[]byte("🍺"), []byte(":beer:")},
	{[]byte("🍻"), []byte(":beers:")},
	{[]byte("🍸"), []byte(":cocktail:")},
	{[]byte("🍹"), []byte(":tropical_drink:")},
	{[]byte("🍷"), []byte(":wine_glass:")},
	{[]byte("🍴"), []byte(":fork_and_knife:")},
	{[]byte("🍕"), []byte(":pizza:")},
	{[]byte("🍔"), []byte(":hamburger:")},
	{[]byte("🍟"), []byte(":fries:")},
	{[]byte("🍗"), []byte(":poultry_leg:")},
	{[]byte("🍖"), []byte(":meat_on_bone:")},
	{[]byte("🍝"), []byte(":spaghetti:")},
	{[]byte("🍛"), []byte(":curry:")},
	{[]byte("🍤"), []byte(":fried_shrimp:")},
	{[]byte("🍱"), []byte(":bento:")},
	{[]byte("🍣"), []byte(":sushi:")},
	{[]byte("🍥"), []byte(":fish_cake:")},
	{[]byte("🍙"), []byte(":rice_ball:")},
	{[]byte("🍘"), []byte(":rice_cracker:")},
	{[]byte("🍚"), []byte(":rice:")},
	{[]byte("🍜"), []byte(":ramen:")},
	{[]byte("🍲"), []byte(":stew:")},
	{[]byte("🍢"), []byte(":oden:")},
	{[]byte("🍡"), []byte(":dango:")},
	{[]byte("🥚"), []byte(":egg:")},
	{[]byte("🍞"), []byte(":bread:")},
	{[]byte("🍩"), []byte(":doughnut:")},
	{[]byte("🍮"), []byte(":custard:")},
	{[]byte("🍦"), []byte(":icecream:")},
	{[]byte("🍨"), []byte(":ice_cream:")},
	{[]byte("🍧"), []byte(":shaved_ice:")},
	{[]byte("🎂"), []byte(":birthday:")},
	{[]byte("🍰"), []byte(":cake:")},
	{[]byte("🍪"), []byte(":cookie:")},
	{[]byte("🍫"), []byte(":chocolate_bar:")},
	{[]byte("🍬"), []byte(":candy:")},
	{[]byte("🍭"), []byte(":lollipop:")},
	{[]byte("🍯"), []byte(":honey_pot:")},
	{[]byte("🍎"), []byte(":apple:")},
	{[]byte("🍏"), []byte(":green_apple:")},
	{[]byte("🍊"), []byte(":tangerine:")},
	{[]byte("🍋"), []byte(":lemon:")},
	{[]byte("🍒"), []byte(":cherries:")},
	{[]byte("🍇"), []byte(":grapes:")},
	{[]byte("🍉"), []byte(":watermelon:")},
	{[]byte("🍓"), []byte(":strawberry:")},
	{[]byte("🍑"), []byte(":peach:")},
	{[]byte("🍈"), []byte(":melon:")},
	{[]byte("🍌"), []byte(":banana:")},
	{[]byte("🍐"), []byte(":pear:")},
	{[]byte("🍍"), []byte(":pineapple:")},
	{[]byte("🍠"), []byte(":sweet_potato:")},
	{[]byte("🍆"), []byte(":eggplant:")},
	{[]byte("🍅"), []byte(":tomato:")},
	{[]byte("🌽"), []byte(":corn:")},
	{[]byte("🏢"), []byte(":office:")},
	{[]byte("🏣"), []byte(":post_office:")},
	{[]byte("🏥"), []byte(":hospital:")},
	{[]byte("🏦"), []byte(":bank:")},
	{[]byte("🏪"), []byte(":convenience_store:")},
	{[]byte("🏩"), []byte(":love_hotel:")},
	{[]byte("🏨"), []byte(":hotel:")},
	{[]byte("💒"), []byte(":wedding:")},
	{[]byte("⛪"), []byte(":church:")},
	{[]byte("🏬"), []byte(":department_store:")},
	{[]byte("🏤"), []byte(":european_post_office:")},
	{[]byte("🌇"), []byte(":city_sunrise:")},
	{[]byte("🌆"), []byte(":city_sunset:")},
	{[]byte("🏯"), []byte(":japanese_castle:")},
	{[]byte("🏰"), []byte(":european_castle:")},
	{[]byte("⛺"), []byte(":tent:")},
	{[]byte("🏭"), []byte(":factory:")},
	{[]byte("🗼"), []byte(":tokyo_tower:")},
	{[]byte("🗾"), []byte(":japan:")},
	{[]byte("🗻"), []byte(":mount_fuji:")},
	{[]byte("🌄"), []byte(":sunrise_over_mountains:")},
	{[]byte("🌅"), []byte(":sunrise:")},
	{[]byte("🌠"), []byte(":stars:")},
	{[]byte("🗽"), []byte(":statue_of_liberty:")},
	{[]byte("🌉"), []byte(":bridge_at_night:")},
	{[]byte("🎠"), []byte(":carousel_horse:")},
	{[]byte("🌈"), []byte(":rainbow:")},
	{[]byte("🎡"), []byte(":ferris_wheel:")},
	{[]byte("⛲"), []byte(":fountain:")},
	{[]byte("🎢"), []byte(":roller_coaster:")},
	{[]byte("🚢"), []byte(":ship:")},
	{[]byte("🚤"), []byte(":speedboat:")},
	{[]byte("⛵"), []byte(":boat:")},
	{[]byte("⛵"), []byte(":sailboat:")},
	{[]byte("🚣"), []byte(":rowboat:")},
	{[]byte("⚓"), []byte(":anchor:")},
	{[]byte("🚀"), []byte(":rocket:")},
	{[]byte("✈️"), []byte(":airplane:")},
	{[]byte("🚁"), []byte(":helicopter:")},
	{[]byte("🚂"), []byte(":steam_locomotive:")},
	{[]byte("🚊"), []byte(":tram:")},
	{[]byte("🚞"), []byte(":mountain_railway:")},
	{[]byte("🚲"), []byte(":bike:")},
	{[]byte("🚡"), []byte(":aerial_tramway:")},
	{[]byte("🚟"), []byte(":suspension_railway:")},
	{[]byte("🚠"), []byte(":mountain_cableway:")},
	{[]byte("🚜"), []byte(":tractor:")},
	{[]byte("🚙"), []byte(":blue_car:")},
	{[]byte("🚘"), []byte(":oncoming_automobile:")},
	{[]byte("🚗"), []byte(":car:")},
	{[]byte("🚗"), []byte(":red_car:")},
	{[]byte("🚕"), []byte(":taxi:")},
	{[]byte("🚖"), []byte(":oncoming_taxi:")},
	{[]byte("🚛"), []byte(":articulated_lorry:")},
	{[]byte("🚌"), []byte(":bus:")},
	{[]byte("🚍"), []byte(":oncoming_bus:")},
	{[]byte("🚨"), []byte(":rotating_light:")},
	{[]byte("🚓"), []byte(":police_car:")},
	{[]byte("🚔"), []byte(":oncoming_police_car:")},
	{[]byte("🚒"), []byte(":fire_engine:")},
	{[]byte("🚑"), []byte(":ambulance:")},
	{[]byte("🚐"), []byte(":minibus:")},
	{[]byte("🚚"), []byte(":truck:")},
	{[]byte("🚋"), []byte(":train:")},
	{[]byte("🚉"), []byte(":station:")},
	{[]byte("🚆"), []byte(":train2:")},
	{[]byte("🚅"), []byte(":bullettrain_front:")},
	{[]byte("🚄"), []byte(":bullettrain_side:")},
	{[]byte("🚈"), []byte(":light_rail:")},
	{[]byte("🚝"), []byte(":monorail:")},
	{[]byte("🚃"), []byte(":railway_car:")},
	{[]byte("🚎"), []byte(":trolleybus:")},
	{[]byte("🎫"), []byte(":ticket:")},
	{[]byte("⛽"), []byte(":fuelpump:")},
	{[]byte("🚦"), []byte(":vertical_traffic_light:")},
	{[]byte("🚥"), []byte(":traffic_light:")},
	{[]byte("⚠️"), []byte(":warning:")},
	{[]byte("🚧"), []byte(":construction:")},
	{[]byte("🔰"), []byte(":beginner:")},
	{[]byte("🏧"), []byte(":atm:")},
	{[]byte("🎰"), []byte(":slot_machine:")},
	{[]byte("🚏"), []byte(":busstop:")},
	{[]byte("💈"), []byte(":barber:")},
	{[]byte("♨️"), []byte(":hotsprings:")},
	{[]byte("🏁"), []byte(":checkered_flag:")},
	{[]byte("🎌"), []byte(":crossed_flags:")},
	{[]byte("🏮"), []byte(":izakaya_lantern:")},
	{[]byte("🗿"), []byte(":moyai:")},
	{[]byte("🎪"), []byte(":circus_tent:")},
	{[]byte("🎭"), []byte(":performing_arts:")},
	{[]byte("📍"), []byte(":round_pushpin:")},
	{[]byte("🚩"), []byte(":triangular_flag_on_post:")},
	{[]byte("🇯🇵"), []byte(":jp:")},
	{[]byte("🇰🇷"), []byte(":kr:")},
	{[]byte("🇨🇳"), []byte(":cn:")},
	{[]byte("🇺🇸"), []byte(":us:")},
	{[]byte("🇫🇷"), []byte(":fr:")},
	{[]byte("🇪🇸"), []byte(":es:")},
	{[]byte("🇮🇹"), []byte(":it:")},
	{[]byte("🇷🇺"), []byte(":ru:")},
	{[]byte("🇬🇧"), []byte(":gb:")},
	{[]byte("🇬🇧"), []byte(":uk:")},
	{[]byte("🇩🇪"), []byte(":de:")},
	{[]byte("4️⃣"), []byte(":four:")},
	{[]byte("5️⃣"), []byte(":five:")},
	{[]byte("6️⃣"), []byte(":six:")},
	{[]byte("7️⃣"), []byte(":seven:")},
	{[]byte("8️⃣"), []byte(":eight:")},
	{[]byte("9️⃣"), []byte(":nine:")},
	{[]byte("🔟"), []byte(":keycap_ten:")},
	{[]byte("🔢"), []byte(":1234:")},
	{[]byte("0️⃣"), []byte(":zero:")},
	{[]byte("#️⃣"), []byte(":hash:")},
	{[]byte("🔣"), []byte(":symbols:")},
	{[]byte("◀️"), []byte(":arrow_backward:")},
	{[]byte("⬇️"), []byte(":arrow_down:")},
	{[]byte("▶️"), []byte(":arrow_forward:")},
	{[]byte("⬅️"), []byte(":arrow_left:")},
	{[]byte("🔠"), []byte(":capital_abcd:")},
	{[]byte("🔡"), []byte(":abcd:")},
	{[]byte("🔤"), []byte(":abc:")},
	{[]byte("↙️"), []byte(":arrow_lower_left:")},
	{[]byte("↘️"), []byte(":arrow_lower_right:")},
	{[]byte("➡️"), []byte(":arrow_right:")},
	{[]byte("⬆️"), []byte(":arrow_up:")},
	{[]byte("↖️"), []byte(":arrow_upper_left:")},
	{[]byte("↗️"), []byte(":arrow_upper_right:")},
	{[]byte("⏬"), []byte(":arrow_double_down:")},
	{[]byte("⏫"), []byte(":arrow_double_up:")},
	{[]byte("🔽"), []byte(":arrow_down_small:")},
	{[]byte("⤵️"), []byte(":arrow_heading_down:")},
	{[]byte("⤴️"), []byte(":arrow_heading_up:")},
	{[]byte("↩️"), []byte(":leftwards_arrow_with_hook:")},
	{[]byte("↪️"), []byte(":arrow_right_hook:")},
	{[]byte("↔️"), []byte(":left_right_arrow:")},
	{[]byte("↕️"), []byte(":arrow_up_down:")},
	{[]byte("🔼"), []byte(":arrow_up_small:")},
	{[]byte("🔃"), []byte(":arrows_clockwise:")},
	{[]byte("🔄"), []byte(":arrows_counterclockwise:")},
	{[]byte("⏪"), []byte(":rewind:")},
	{[]byte("⏩"), []byte(":fast_forward:")},
	{[]byte("ℹ️"), []byte(":information_source:")},
	{[]byte("🆗"), []byte(":ok:")},
	{[]byte("🔀"), []byte(":twisted_rightwards_arrows:")},
	{[]byte("🔁"), []byte(":repeat:")},
	{[]byte("🔂"), []byte(":repeat_one:")},
	{[]byte("🆕"), []byte(":new:")},
	{[]byte("🔝"), []byte(":top:")},
	{[]byte("🆙"), []byte(":up:")},
	{[]byte("🆒"), []byte(":cool:")},
	{[]byte("🆓"), []byte(":free:")},
	{[]byte("🆖"), []byte(":ng:")},
	{[]byte("🎦"), []byte(":cinema:")},
	{[]byte("🈁"), []byte(":koko:")},
	{[]byte("📶"), []byte(":signal_strength:")},
	{[]byte("🈹"), []byte(":u5272:")},
	{[]byte("🈴"), []byte(":u5408:")},
	{[]byte("🈺"), []byte(":u55b6:")},
	{[]byte("🈯"), []byte(":u6307:")},
	{[]byte("🈷️"), []byte(":u6708:")},
	{[]byte("🈶"), []byte(":u6709:")},
	{[]byte("🈵"), []byte(":u6e80:")},
	{[]byte("🈚"), []byte(":u7121:")},
	{[]byte("🈸"), []byte(":u7533:")},
	{[]byte("🈳"), []byte(":u7a7a:")},
	{[]byte("🈲"), []byte(":u7981:")},
	{[]byte("🈂️"), []byte(":sa:")},
	{[]byte("🚻"), []byte(":restroom:")},
	{[]byte("🚹"), []byte(":mens:")},
	{[]byte("🚺"), []byte(":womens:")},
	{[]byte("🚼"), []byte(":baby_symbol:")},
	{[]byte("🚭"), []byte(":no_smoking:")},
	{[]byte("🅿️"), []byte(":parking:")},
	{[]byte("♿"), []byte(":wheelchair:")},
	{[]byte("🚇"), []byte(":metro:")},
	{[]byte("🛄"), []byte(":baggage_claim:")},
	{[]byte("🉑"), []byte(":accept:")},
	{[]byte("🚾"), []byte(":wc:")},
	{[]byte("🚰"), []byte(":potable_water:")},
	{[]byte("🚮"), []byte(":put_litter_in_its_place:")},
	{[]byte("㊙️"), []byte(":secret:")},
	{[]byte("㊗️"), []byte(":congratulations:")},
	{[]byte("Ⓜ️"), []byte(":m:")},
	{[]byte("🛂"), []byte(":passport_control:")},
	{[]byte("🛅"), []byte(":left_luggage:")},
	{[]byte("🛃"), []byte(":customs:")},
	{[]byte("🉐"), []byte(":ideograph_advantage:")},
	{[]byte("🆑"), []byte(":cl:")},
	{[]byte("🆘"), []byte(":sos:")},
	{[]byte("🆔"), []byte(":id:")},
	{[]byte("🚫"), []byte(":no_entry_sign:")},
	{[]byte("🔞"), []byte(":underage:")},
	{[]byte("📵"), []byte(":no_mobile_phones:")},
	{[]byte("🚯"), []byte(":do_not_litter:")},
	{[]byte("🚱"), []byte(":non-potable_water:")},
	{[]byte("🚳"), []byte(":no_bicycles:")},
	{[]byte("🚷"), []byte(":no_pedestrians:")},
	{[]byte("🚸"), []byte(":children_crossing:")},
	{[]byte("⛔"), []byte(":no_entry:")},
	{[]byte("✳️"), []byte(":eight_spoked_asterisk:")},
	{[]byte("✴️"), []byte(":eight_pointed_black_star:")},
	{[]byte("💟"), []byte(":heart_decoration:")},
	{[]byte("🆚"), []byte(":vs:")},
	{[]byte("📳"), []byte(":vibration_mode:")},
	{[]byte("📴"), []byte(":mobile_phone_off:")},
	{[]byte("💹"), []byte(":chart:")},
	{[]byte("💱"), []byte(":currency_exchange:")},
	{[]byte("♈"), []byte(":aries:")},
	{[]byte("♉"), []byte(":taurus:")},
	{[]byte("♊"), []byte(":gemini:")},
	{[]byte("♋"), []byte(":cancer:")},
	{[]byte("♌"), []byte(":leo:")},
	{[]byte("♍"), []byte(":virgo:")},
	{[]byte("♎"), []byte(":libra:")},
	{[]byte("♏"), []byte(":scorpius:")},
	{[]byte("♐"), []byte(":sagittarius:")},
	{[]byte("♑"), []byte(":capricorn:")},
	{[]byte("♒"), []byte(":aquarius:")},
	{[]byte("♓"), []byte(":pisces:")},
	{[]byte("⛎"), []byte(":ophiuchus:")},
	{[]byte("🔯"), []byte(":six_pointed_star:")},
	{[]byte("❎"), []byte(":negative_squared_cross_mark:")},
	{[]byte("🅰️"), []byte(":a:")},
	{[]byte("🅱️"), []byte(":b:")},
	{[]byte("🆎"), []byte(":ab:")},
	{[]byte("🅾️"), []byte(":o2:")},
	{[]byte("💠"), []byte(":diamond_shape_with_a_dot_inside:")},
	{[]byte("♻️"), []byte(":recycle:")},
	{[]byte("🔚"), []byte(":end:")},
	{[]byte("🔛"), []byte(":on:")},
	{[]byte("🔜"), []byte(":soon:")},
	{[]byte("🕐"), []byte(":clock1:")},
	{[]byte("🕜"), []byte(":clock130:")},
	{[]byte("🕙"), []byte(":clock10:")},
	{[]byte("🕥"), []byte(":clock1030:")},
	{[]byte("🕚"), []byte(":clock11:")},
	{[]byte("🕦"), []byte(":clock1130:")},
	{[]byte("🕛"), []byte(":clock12:")},
	{[]byte("🕧"), []byte(":clock1230:")},
	{[]byte("🕑"), []byte(":clock2:")},
	{[]byte("🕝"), []byte(":clock230:")},
	{[]byte("🕒"), []byte(":clock3:")},
	{[]byte("🕞"), []byte(":clock330:")},
	{[]byte("🕓"), []byte(":clock4:")},
	{[]byte("🕟"), []byte(":clock430:")},
	{[]byte("🕔"), []byte(":clock5:")},
	{[]byte("🕠"), []byte(":clock530:")},
	{[]byte("🕕"), []byte(":clock6:")},
	{[]byte("🕡"), []byte(":clock630:")},
	{[]byte("🕖"), []byte(":clock7:")},
	{[]byte("🕢"), []byte(":clock730:")},
	{[]byte("🕗"), []byte(":clock8:")},
	{[]byte("🕣"), []byte(":clock830:")},
	{[]byte("🕘"), []byte(":clock9:")},
	{[]byte("🕤"), []byte(":clock930:")},
	{[]byte("💲"), []byte(":heavy_dollar_sign:")},
	{[]byte("©️"), []byte(":copyright:")},
	{[]byte("®️"), []byte(":registered:")},
	{[]byte("™️"), []byte(":tm:")},
	{[]byte("❌"), []byte(":x:")},
	{[]byte("❗"), []byte(":heavy_exclamation_mark:")},
	{[]byte("‼️"), []byte(":bangbang:")},
	{[]byte("⁉️"), []byte(":interrobang:")},
	{[]byte("⭕"), []byte(":o:")},
	{[]byte("✖️"), []byte(":heavy_multiplication_x:")},
	{[]byte("➕"), []byte(":heavy_plus_sign:")},
	{[]byte("➖"), []byte(":heavy_minus_sign:")},
	{[]byte("➗"), []byte(":heavy_division_sign:")},
	{[]byte("💮"), []byte(":white_flower:")},
	{[]byte("💯"), []byte(":100:")},
	{[]byte("✔️"), []byte(":heavy_check_mark:")},
	{[]byte("☑️"), []byte(":ballot_box_with_check:")},
	{[]byte("🔘"), []byte(":radio_button:")},
	{[]byte("🔗"), []byte(":link:")},
	{[]byte("➰"), []byte(":curly_loop:")},
	{[]byte("〰️"), []byte(":wavy_dash:")},
	{[]byte("〽️"), []byte(":part_alternation_mark:")},
	{[]byte("🔱"), []byte(":trident:")},
	{[]byte("✅"), []byte(":white_check_mark:")},
	{[]byte("🔲"), []byte(":black_square_button:")},
	{[]byte("🔳"), []byte(":white_square_button:")},
	{[]byte("⚫"), []byte(":black_circle:")},
	{[]byte("⚪"), []byte(":white_circle:")},
	{[]byte("🔴"), []byte(":red_circle:")},
	{[]byte("🔵"), []byte(":large_blue_circle:")},
	{[]byte("🔷"), []byte(":large_blue_diamond:")},
	{[]byte("🔶"), []byte(":large_orange_diamond:")},
	{[]byte("🔹"), []byte(":small_blue_diamond:")},
	{[]byte("🔸"), []byte(":small_orange_diamond:")},
	{[]byte("🔺"), []byte(":small_red_triangle:")},
	{[]byte("🔻"), []byte(":small_red_triangle_down:")},
}
