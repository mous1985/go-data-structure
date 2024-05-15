# Tries
 A Trie, also known as a prefix tree, is a type of search tree used in computer science for storing a collection of strings. It has a number of uses in data structures and algorithms, including:

Autocomplete: Tries can be used to suggest the rest of a word when the user has typed in the first few letters.
Spell check: Tries can be used to suggest corrections for misspelled words.
IP routing: Tries can be used to store and search IP addresses.
In a Trie, each node represents a character of a string and each path down the tree can represent a complete string. The root node is usually associated with the empty string, and the children of a node have labels that are all different and correspond to the characters that might follow the string associated with their parent node in the set of strings the Trie is representing.

Here's an example of what a Trie might look like when storing the words "try", "trie", and "tried":

  root
   |
   t
   |
   r
   |
   y - i - e - d
   |
   e
   In this Trie, the path from the root to the 'y' node represents the string "try", the path to the 'e' node represents "trie", and the path to the 'd' node represents "tried".