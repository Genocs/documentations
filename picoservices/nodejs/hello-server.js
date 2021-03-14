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

app.get('/ping', (_req, res) => {
    res.send('pong');
});

app.post('/post_my_data', express.json(), (req, res) => {
    const data = req.body.data;
    console.log(req.body);

    var result = '';

    if (data) {
        const id = data.id;
        result = "Got a body with ID: " + id;
    } else {
        result = "The body data is empty";
    }
    console.log(result);
    res.status(200).send(result);
});

app.listen(port, () => console.log(`Node App listening on port ${port}!`));