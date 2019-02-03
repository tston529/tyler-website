from css_html_js_minify import process_single_js_file, process_single_css_file
import os
from rjsmin import jsmin

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

for d in dirs:
	files = os.listdir(d)
	for file in files:
		if ".min." in file:
			print("Removing old copy of " + d+"/"+file)
			os.remove(d+"/"+file)
	files = os.listdir(d)
	for file in files:
		fn = file.split(".")
		minify_request(file, d)