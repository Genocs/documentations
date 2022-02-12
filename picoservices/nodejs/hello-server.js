// ------------------------------------------------------------
// Giovanni Emanuele Nocco [Genocs].
// Licensed under the MIT License.
// ------------------------------------------------------------

const express = require('express');

const app = express();
app.use(express.json());
const port = 5300;

app.get('/', (_req, res) => {
    console.log('Called home!!!');
    res.send('Hello from Node.js!!!');
});

app.get('/hc', (_req, res) => {
    res.send('healthy');
});

app.post('/order/submit', express.json(), (req, res) => {
    //console.log(req.body);
    const data = req.body.data;
    var result = {};

    if (data) {
        result = { id: data.id, description: data.description, status: "Accepted" };
    } else {
        result = "The body data is empty";
    }
    //console.log(result);
    res.status(200).send(result);
});

app.listen(port, () => console.log(`Node App listening on port ${port}!`));