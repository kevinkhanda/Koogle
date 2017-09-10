# Koogle
A Search Engine which in perspective will destroy Google. (At least because Google developers may die of laughter)

## Getting started
Project is developed using Go programming language.

### Design decisions
Indexing in Koogle is based on inverted index with term frequency idea.
It means that the documents in result of a query will be sorted with respect
to term frequencies. 

#### Indexing
Here is the idea about how the indexing happens:
* First, corpus is divided into a separate documents
* After the separation, documents are tokenized
* Koogle is storing a map where keys are tokens and values are lists
 of documents IDs
* After that the file "invertedIndex" is generated (you can check it in 
"index" directory)
    * As you can see, single entry in this file is look like this
    ```
        "term" -> <docId1:termFrequency1>...<docIdN:termFrequencyN> 
    ```
    * Postings list is sorted by descending by termFrequency
* After the indexes file was generated, Koogle generates the "stemmingData"
file. The idea is to map word roots with list of words.
    * In "stemmingData" file entries are represented as:
    ```
        "word root" -> <term1:termIndexInIndexesFile1>...<termN:termIndexInIndexesFileN>
    ```
    * This design decision was done in order to have a faster lookup for
    desired term in indexes file

#### Searching
The postings for several terms in user query are merged. So, in result you
will see a queries, in which all of the query terms are presented.

Also you will not see the whole document with content. Only document title
and number will be shown. 

Here is how the searching is working:
* First, after Koogle receives the user input, it splits it on terms
* For each term then, it is looking for a root (stemming the term)
* After that a lookup is made in a "stemmingData" file
* When the required root is found, we are also looking for a term and its
index in "invertedIndex" file
* The next step is to find an term according to its index and retrieve 
a list of postings in which this term appears.
* The last step is to retrieve postings titles and show them to user


## Requirements
Apple macOS: Install [Go](https://storage.googleapis.com/golang/go1.9.darwin-amd64.pkg)

Microsoft Windows: Install [Go](https://storage.googleapis.com/golang/go1.9.windows-amd64.msi)

Linux: Install [Go](https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz)

### Installation
After you have installed Go language from the link provided before, do the following in Terminal:

**Works for Linux and macOS systems**
```
mkdir ~/go/src/github.com/caneroj1/; 
git clone https://github.com/caneroj1/stemmer.git ~/go/src/github.com/caneroj1/
mkdir ~/go/src/github.com/kevinkhanda/; 
git clone https://github.com/KKhanda/Koogle.git ~/go/src/github.com/kevinkhanda/
```
**Windows Users**


### Building and running the application
Application is now presented only as a CLI (will be extended later).
It is OK to see a lot of error outputs. This means that some strings 
were not matched by regular expressions.

#### Running application with building
```
cd ~/go/src/github.com/kevinkhanda/koogle/main/
go build
./main
```

#### Running the application without build (build file is already included)
```
cd ~/go/src/github.com/kevinkhanda/koogle/main/
./main
```

## How to use the application 
You will asked by a command prompt to type a query. So, this is what
you should do. You can type any kind of query (number of words is not
limited), however, in some cases you won't receive a result.

Here are the examples of queries:

![Alt text](screenshots/koogle_launch.png?raw=true "Koogle greetings!")
![Alt text](screenshots/koogle_prompt.png?raw=true "Koogle prompt")
![Alt text](screenshots/koogle_single_result.png?raw=true "Simple query")
![Alt text](screenshots/koogle_several_results.png?raw=true "Several results")
![Alt text](screenshots/not_found_koogle.png?raw=true "Error example")
![Alt text](screenshots/Very_complex_query.png?raw=true "Complex query")
