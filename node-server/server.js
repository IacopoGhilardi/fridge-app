const express = require('express');
const passport = require('passport');
const mongoose = require('mongoose');
const Fridge = require('./models/fridge');
const cors = require('cors')
const bodyParser = require('body-parser');
const fridge = require('./models/fridge');

// const passportVar = passport();
const app = express();
app.use(cors());
app.use(bodyParser.json());

app.listen(3000, () => {
    console.log('listening on port 3000');
})
mongoose.connect('mongodb://localhost:27017/the-fridge');

app.get('/', function (req, res) {
    res.redirect('http://localhost:8080/')
});

app.post('/fridge/upsert/:fridgeId', cors(), async function (req, res) {
  console.log(`Upserting fridge: ${req.params.fridgeId}`);

  try {
    await Fridge.findOneAndUpdate({_id: `${req.params.fridgeId}`}, {$set:{food:req.body.fridge}})

    res.send({
      'Success': 'Fridge updated'
    });
  } catch (error) {
    console.log(error);
    res.send({
      'error': error
    });
  }
});

app.get('/fridge', cors(), async function (req, res) {
    console.log('chiamata get');

    try {
      let fridgeItems = await Fridge.find();
      res.send({
        'fridgeItems': fridgeItems
      })
    } catch (error) {
      console.log(error);
      res.send({
        'error': error,
      })
    }
});