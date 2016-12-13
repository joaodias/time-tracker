echo "- Removing public directory"
rm -rf public
echo "- Making public directory"
mkdir public
echo "- Generating public content (Webpacking)"
webpack
echo "- Copying stuff"
cp index.html public
cp app.css public
echo "- Running surge"
cd public
surge
echo "Success!"