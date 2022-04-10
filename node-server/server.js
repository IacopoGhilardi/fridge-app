const express = require('express');
const passport = require('passport');
const mongoose = require('mongoose');
const bcrypt = require('bcryptjs');
const jwt = require("jsonwebtoken");
const Fridge = require('./models/fridge');
const cors = require('cors')
const bodyParser = require('body-parser');
const User = require('./models/user');
const authMiddleware = require("./middleware/auth");
const app = express();
require("dotenv").config();
app.use(cors());
app.use(bodyParser.json());

app.listen(3000, () => {
    console.log('listening on port 3000');
})
mongoose.connect('mongodb://localhost:27017/the-fridge');

app.post('/register', async (req, res) => {
  try {
    // Get user input
    const { first_name, last_name, email, password } = req.body;

    //validate
    if (!(email && password && first_name && last_name)) {
      res.status(400).send("All input is required");
    }

    const oldUser = await User.findOne({email});

    if (oldUser) {
      return res.status(409).send("User Already Exist. Pleas login");
    }
    
    encryptedPassword = await bcrypt.hash(password, 10);

    const user = await User.create({
      first_name,
      last_name,
      email: email,
      password: encryptedPassword
    });

    const token = jwt.sign(
      { user_id: user._id, email },
      process.env.TOKEN_KEY,
      {
        expiresIn: "2h",
      }
    )

    console.log('token', token);
    user.token = token;

    res.status(201).json(user);

  } catch (error) {
    console.log(error);
    res.status(500).json(error);
  }
});

app.post('/login', async (req, res) => {
  try {

    const {email, password} = req.body;

    // Validate user input
    if (!(email && password)) {
      res.status(400).send("All input is required");
    }

    // Validate if user exist in our database
    const user = await User.findOne({ email });

    if (user && (await bcrypt.compare(password, user.password))) {
      console.log('Password confermata');
      //create token
      const token = jwt.sign(
        { user_id: user._id, email },
        process.env.TOKEN_KEY,
        {
          expiresIn: "2h",
        }
      );

      user.token = token;

    } else {
      console.log('Password doesn\'t match');
      res.status(400).send("Password does not match");
    }

    console.log('arrivi qua?');
    // console.log('token', token);
    // save user token
    

    console.log('ora?');

    

    // user
    res.status(200).json(user);

  } catch (error) {
    console.log(error);
    res.send({
      'error': error
    });
  }
});


app.get('/', (req, res) => {
    res.redirect('http://localhost:8080/')
});

app.post('/fridge/upsert/:fridgeId', authMiddleware, cors(), async (req, res) => {
  console.log('upsert fridge');

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

app.get('/fridge',  authMiddleware, cors(), async (req, res) => {
    try {
      let fridgeItems = await Fridge.find();

      res.send({
        "foodList": fridgeItems
      })
    } catch (error) {
      console.log(error);
      res.send({
        'error': error,
      })
    }
});