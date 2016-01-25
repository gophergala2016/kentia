# Kentia easy dress

When you don’t have idea about fashion and you waste a lot of time trying to get a nice outfit to use, think in an application that helps you to combine the color of your clothes, It will save your life and time.

We share the necessity and we took on the task and we created this application to solve this problem.

Kentia is a website that helps you to combine your clothes using a genetic algorithm, we use some characteristics for your clothes so we can classify them.

The name “Kentia” means “dress me” in Nahuatl.

## Installation

On server you have to compile the project using

```go install kentia```

>Make sure you have already set your environment variable GOPATH

You have to restore the database of the directory in Kentia/mongodb

```mongorestore mongodb/dump```

If you want to run the server just run the binary on the root directory of the project , the server needs some static resources (public,templates,etc).

After you run the binary the server will use port `:3000` for listening requests.

## How does it work

To start using Kentia the user must to sign up into the system, giving his name, email, gender and password. When the user registers on Kentia he will be able to upload pictures and information about his clothes in order to have an online wardrobe.

Once that the user has his online wardrobe, Kentia will be able to suggest different combinations.
To know how to combine the clothes, we did a little research about how to combine colors properly, after that we make a genetic algorithm to determine the best color combinations that satisfy in a good way this theory considering the clothes that the user have.

The combine method is defined with some concepts like complementary, analog and monochromatic colors, those combinations are very important in the evaluation functions of genetic algorithm. Another parameter is the type of clothes (footwear,shirts,jeans and jackets). Using the genetic algorithm, we try to get a high aptitude taking the combinations of the colors.

When you are trying to get a combination the website will give the 3 higher combinations of your wardrobe.

## Screenshots
![Combination 1](http://i.imgur.com/40ftS1s.png)
![Combination 2](http://i.imgur.com/rLdLMzT.png)

## TODO

* Take more properties to generate a better custom combination to the user
* Image analyzing to get the color and type of the clothes
* Combinate textures like squares or bars and not only solid colors
* Integrate the application whit social networks
* Implement the favorite combination to generate a combination accord your preferences
* Show to the user the last date when he uses that clothes

## Team

* [Mario Carrillo](https://github.com/fiberto)
* [Juan Torres](https://github.com/JuanTorr)
* [Noe Eustaquio](https://github.com/nedorowsky)
* [Carlos Granados](https://github.com/remnanttime)
