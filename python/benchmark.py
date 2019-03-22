from random import randint
import time

from pyexcelerate import Workbook

ROWS = 100000
COLS = 50

def generate_data():
	data = [[0] * COLS] * ROWS
	for row in range(ROWS):
		for col in range(COLS):
			data[row][col] = randint(1, 10000)
	return data

def test():
	begin_timestamp = time.time()
	wb = Workbook()
	wb.new_sheet("sheet name", data=generate_data())
	wb.save("output/python.xlsx")
	end_timestamp = time.time()

	print("Python: Writing 10000x50 cells of data takes %f seconds" % (end_timestamp - begin_timestamp))


if __name__ == '__main__':
	test()
