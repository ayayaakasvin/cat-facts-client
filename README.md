# cat-facts-client
A small Go client for the MeowFacts public API.
Exposes a Client type with a FunFact method that returns a random cat fact. 
Uses a configurable http.Client with a sensible default timeout.
Lightweight and suitable for CLI tools or small services.
Minimal dependencies and easy to embed in other projects.
Huge props to MeowFact!