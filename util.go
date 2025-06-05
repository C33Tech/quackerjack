package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Stop Words
var stopWords = map[string]bool{"a": true, "able": true, "about": true, "above": true, "abroad": true, "according": true, "accordingly": true, "across": true, "actually": true, "adj": true, "after": true, "afterwards": true, "again": true, "against": true, "ago": true, "ahead": true, "ain't": true, "all": true, "allow": true, "allows": true, "almost": true, "alone": true, "along": true, "alongside": true, "already": true, "also": true, "although": true, "always": true, "am": true, "amid": true, "amidst": true, "among": true, "amongst": true, "an": true, "and": true, "another": true, "any": true, "anybody": true, "anyhow": true, "anyone": true, "anything": true, "anyway": true, "anyways": true, "anywhere": true, "apart": true, "appear": true, "appreciate": true, "appropriate": true, "are": true, "aren't": true, "around": true, "as": true, "a's": true, "aside": true, "ask": true, "asking": true, "associated": true, "at": true, "available": true, "away": true, "awfully": true, "b": true, "back": true, "backward": true, "backwards": true, "be": true, "became": true, "because": true, "become": true, "becomes": true, "becoming": true, "been": true, "before": true, "beforehand": true, "begin": true, "behind": true, "being": true, "believe": true, "below": true, "beside": true, "besides": true, "best": true, "better": true, "between": true, "beyond": true, "both": true, "brief": true, "but": true, "by": true, "c": true, "came": true, "can": true, "cannot": true, "cant": true, "can't": true, "caption": true, "cause": true, "causes": true, "certain": true, "certainly": true, "changes": true, "clearly": true, "c'mon": true, "co": true, "co.": true, "com": true, "come": true, "comes": true, "concerning": true, "consequently": true, "consider": true, "considering": true, "contain": true, "containing": true, "contains": true, "corresponding": true, "could": true, "couldn't": true, "course": true, "c's": true, "currently": true, "d": true, "dare": true, "daren't": true, "definitely": true, "described": true, "despite": true, "did": true, "didn't": true, "different": true, "directly": true, "do": true, "does": true, "doesn't": true, "doing": true, "done": true, "don't": true, "down": true, "downwards": true, "during": true, "e": true, "each": true, "edu": true, "eg": true, "eight": true, "eighty": true, "either": true, "else": true, "elsewhere": true, "end": true, "ending": true, "enough": true, "entirely": true, "especially": true, "et": true, "etc": true, "even": true, "ever": true, "evermore": true, "every": true, "everybody": true, "everyone": true, "everything": true, "everywhere": true, "ex": true, "exactly": true, "example": true, "except": true, "f": true, "fairly": true, "far": true, "farther": true, "few": true, "fewer": true, "fifth": true, "first": true, "five": true, "followed": true, "following": true, "follows": true, "for": true, "forever": true, "former": true, "formerly": true, "forth": true, "forward": true, "found": true, "four": true, "from": true, "further": true, "furthermore": true, "g": true, "get": true, "gets": true, "getting": true, "given": true, "gives": true, "go": true, "goes": true, "going": true, "gone": true, "got": true, "gotten": true, "greetings": true, "h": true, "had": true, "hadn't": true, "half": true, "happens": true, "hardly": true, "has": true, "hasn't": true, "have": true, "haven't": true, "having": true, "he": true, "he'd": true, "he'll": true, "hello": true, "help": true, "hence": true, "her": true, "here": true, "hereafter": true, "hereby": true, "herein": true, "here's": true, "hereupon": true, "hers": true, "herself": true, "he's": true, "hi": true, "him": true, "himself": true, "his": true, "hither": true, "hopefully": true, "how": true, "howbeit": true, "however": true, "hundred": true, "i": true, "i'd": true, "ie": true, "if": true, "ignored": true, "i'll": true, "i'm": true, "immediate": true, "in": true, "inasmuch": true, "inc": true, "inc.": true, "indeed": true, "indicate": true, "indicated": true, "indicates": true, "inner": true, "inside": true, "insofar": true, "instead": true, "into": true, "inward": true, "is": true, "isn't": true, "it": true, "it'd": true, "it'll": true, "its": true, "it's": true, "itself": true, "i've": true, "j": true, "just": true, "k": true, "keep": true, "keeps": true, "kept": true, "know": true, "known": true, "knows": true, "l": true, "last": true, "lately": true, "later": true, "latter": true, "latterly": true, "least": true, "less": true, "lest": true, "let": true, "let's": true, "like": true, "liked": true, "likely": true, "likewise": true, "little": true, "look": true, "looking": true, "looks": true, "low": true, "lower": true, "ltd": true, "m": true, "made": true, "mainly": true, "make": true, "makes": true, "many": true, "may": true, "maybe": true, "mayn't": true, "me": true, "mean": true, "meantime": true, "meanwhile": true, "merely": true, "might": true, "mightn't": true, "mine": true, "minus": true, "miss": true, "more": true, "moreover": true, "most": true, "mostly": true, "mr": true, "mrs": true, "much": true, "must": true, "mustn't": true, "my": true, "myself": true, "n": true, "name": true, "namely": true, "nd": true, "near": true, "nearly": true, "necessary": true, "need": true, "needn't": true, "needs": true, "neither": true, "never": true, "neverf": true, "neverless": true, "nevertheless": true, "new": true, "next": true, "nine": true, "ninety": true, "no": true, "nobody": true, "non": true, "none": true, "nonetheless": true, "noone": true, "no-one": true, "nor": true, "normally": true, "not": true, "nothing": true, "notwithstanding": true, "novel": true, "now": true, "nowhere": true, "o": true, "obviously": true, "of": true, "off": true, "often": true, "oh": true, "ok": true, "okay": true, "old": true, "on": true, "once": true, "one": true, "ones": true, "one's": true, "only": true, "onto": true, "opposite": true, "or": true, "other": true, "others": true, "otherwise": true, "ought": true, "oughtn't": true, "our": true, "ours": true, "ourselves": true, "out": true, "outside": true, "over": true, "overall": true, "own": true, "p": true, "particular": true, "particularly": true, "past": true, "per": true, "perhaps": true, "placed": true, "please": true, "plus": true, "possible": true, "presumably": true, "probably": true, "provided": true, "provides": true, "q": true, "que": true, "quite": true, "qv": true, "r": true, "rather": true, "rd": true, "re": true, "really": true, "reasonably": true, "recent": true, "recently": true, "regarding": true, "regardless": true, "regards": true, "relatively": true, "respectively": true, "right": true, "round": true, "s": true, "said": true, "same": true, "saw": true, "say": true, "saying": true, "says": true, "second": true, "secondly": true, "see": true, "seeing": true, "seem": true, "seemed": true, "seeming": true, "seems": true, "seen": true, "self": true, "selves": true, "sensible": true, "sent": true, "serious": true, "seriously": true, "seven": true, "several": true, "shall": true, "shan't": true, "she": true, "she'd": true, "she'll": true, "she's": true, "should": true, "shouldn't": true, "since": true, "six": true, "so": true, "some": true, "somebody": true, "someday": true, "somehow": true, "someone": true, "something": true, "sometime": true, "sometimes": true, "somewhat": true, "somewhere": true, "soon": true, "sorry": true, "specified": true, "specify": true, "specifying": true, "still": true, "sub": true, "such": true, "sup": true, "sure": true, "t": true, "take": true, "taken": true, "taking": true, "tell": true, "tends": true, "th": true, "than": true, "thank": true, "thanks": true, "thanx": true, "that": true, "that'll": true, "thats": true, "that's": true, "that've": true, "the": true, "their": true, "theirs": true, "them": true, "themselves": true, "then": true, "thence": true, "there": true, "thereafter": true, "thereby": true, "there'd": true, "therefore": true, "therein": true, "there'll": true, "there're": true, "theres": true, "there's": true, "thereupon": true, "there've": true, "these": true, "they": true, "they'd": true, "they'll": true, "they're": true, "they've": true, "thing": true, "things": true, "think": true, "third": true, "thirty": true, "this": true, "thorough": true, "thoroughly": true, "those": true, "though": true, "three": true, "through": true, "throughout": true, "thru": true, "thus": true, "till": true, "to": true, "together": true, "too": true, "took": true, "toward": true, "towards": true, "tried": true, "tries": true, "truly": true, "try": true, "trying": true, "t's": true, "twice": true, "two": true, "u": true, "un": true, "under": true, "underneath": true, "undoing": true, "unfortunately": true, "unless": true, "unlike": true, "unlikely": true, "until": true, "unto": true, "up": true, "upon": true, "upwards": true, "us": true, "use": true, "used": true, "useful": true, "uses": true, "using": true, "usually": true, "v": true, "value": true, "various": true, "versus": true, "very": true, "via": true, "viz": true, "vs": true, "w": true, "want": true, "wants": true, "was": true, "wasn't": true, "way": true, "we": true, "we'd": true, "welcome": true, "well": true, "we'll": true, "went": true, "were": true, "we're": true, "weren't": true, "we've": true, "what": true, "whatever": true, "what'll": true, "what's": true, "what've": true, "when": true, "whence": true, "whenever": true, "where": true, "whereafter": true, "whereas": true, "whereby": true, "wherein": true, "where's": true, "whereupon": true, "wherever": true, "whether": true, "which": true, "whichever": true, "while": true, "whilst": true, "whither": true, "who": true, "who'd": true, "whoever": true, "whole": true, "who'll": true, "whom": true, "whomever": true, "who's": true, "whose": true, "why": true, "will": true, "willing": true, "wish": true, "with": true, "within": true, "without": true, "wonder": true, "won't": true, "would": true, "wouldn't": true, "x": true, "y": true, "yes": true, "yet": true, "you": true, "you'd": true, "you'll": true, "your": true, "you're": true, "yours": true, "yourself": true, "yourselves": true, "you've": true, "z": true, "zero": true}

func GetWords(text string) []string {
	words := regexp.MustCompile("\\w+")
	return words.FindAllString(text, -1)
}

func IsStopWord(needle string) bool {
	if stopWords[strings.ToLower(needle)] {
		return true
	}

	return false
}

// LoadStopWords reads stop words from a single file path into the stopWords map.
func LoadStopWords(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stopWords[scanner.Text()] = true
	}
}

// Sorting
// via: https://gist.github.com/ikbear/4038654

type sortedMap struct {
	m map[string]uint64
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func SortedKeys(m map[string]uint64) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}
