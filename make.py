import os
import sys
from rjsmin import jsmin
from css_html_js_minify import process_single_css_file

dirs = ["./scripts", "./styles"]

def minify_request(base_filename, directory):
	input_filename = directory+'/'+base_filename
	if ".js" in input_filename:
		read_data = ""
		with open(input_filename, 'r') as inf:
			read_data = inf.read()
		minified = jsmin(read_data)
		split_filename = base_filename.split(".")
		split_filename.insert(len(split_filename)-1, "min")
		split_filename = ".".join(split_filename)
		print(split_filename)
		with open(directory + '/' + split_filename, "w+") as ouf:
			ouf.write(minified)
		return True
	elif ".css" in input_filename:
		process_single_css_file(input_filename, overwrite=False)
		return True
	else:
		return False

def purge_min(directory):
    files = os.listdir(directory)
    for file in files:
        if ".min." in file:
            print("Removing old copy of " + directory+"/"+file)
            os.remove(directory+"/"+file)

def minify():
    for d in dirs:
        purge_min(d)
        files = os.listdir(d)
        for file in files:
            fn = file.split(".")
            minify_request(file, d)

def git():
        print("Calling 'git stage .'")
        os.system("git stage .")
        print("Staged all changes.")
        os.system("git commit")
        inp = input("Would you like to push? (Y/n): ")
        if inp.lower() == 'y':
            os.system("git push")

def main():
    do_git = False
    if len(sys.argv) > 1:
        do_git = True if sys.argv[1] == "push" else False

    minify()
    if do_git:
        git()

if __name__ == "__main__":
    main()
