echo "Enter the name of the files you want to push"
read files
git add $files
echo "Enter the message for the commit"
read message
git config --global user.email cherifmbaye02@gmail.com
git config --global user.name TanakAiko
git commit -m "$message"
git push