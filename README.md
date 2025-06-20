# go-ponteiros-aula

## 📘 Ponteiros em Go

### 📌 1. O que é um ponteiro?

Imagine que a memória do seu computador é como uma grande estante com gavetas numeradas. 

Cada gaveta pode guardar um valor, e cada uma tem seu número (ou seja, seu endereço na memória).

Um ponteiro é como uma indicação com um número que indica em qual gaveta onde está guardado um valor.

### 🧠 2. Por que usar ponteiros?

- Evitar cópias desnecessárias (pass by reference)
- Alterar valores fora de uma função
- Melhorar performance ao lidar com estruturas grandes

### 🧩 3. Funções com ponteiros

- **Sem** ponteiro: o valor é **copiado** e não muda fora da função.
- **Com** ponteiro: o valor é **alterado** diretamente e muda fora da função.