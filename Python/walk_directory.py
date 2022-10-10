import os
import sys

dir_to_walk = sys.argv[1]

for current_dir, folders, files in os.walk(dir_to_walk):
    # we do not have to worry about '.' and '..' path
    for folder in files:
        print('File: ' + os.path.join(current_dir, folder))
    for folder in folders:
        print('Folder: ' + os.path.join(current_dir, folder))
