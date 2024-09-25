{
    const array_client = [
        {
            name: "John",
            age: 30
        },
        {
            name: "Jane",
            age: 25
        },
        {
            name: "Joe",
            age: 40
        },
        {
            name: "Jill",
            age: 35
        }
    ];

    // Transformando os dados em JSON
    const jsonClient = JSON.stringify(array_client);
    console.log(jsonClient);

    // Transformando JSON em dados
    const client = JSON.parse(jsonClient);
    console.log(client);

    // Imprimindo o nome dos clientes
    console.log("Imprimindo o nome dos clientes: For of");
    for (const value of client) {
        console.log(value.name)
    }

    // Imprimindo o nome dos clientes
    const allNames = client.map((x: { name: string }) => x.name);
    console.log(allNames)

    // Filtrando os clientes maiores de 35 anos
    const filterAge = client.filter((x: { age: number; }) => x.age >= 35);
    console.log(filterAge)

}