require 'axlsx'

ROWS = 10000
COLS = 50

def generate_test_data()
    data = []
    0.upto(ROWS - 1) do |row|
        0.upto(COLS - 1) do |col|
            data[row] ||= []
            data[row][col] = row + col
        end
    end
    data
end

def test()
    test_data = generate_test_data()

    begin_timestamp = Time.now().to_f

    p = Axlsx::Package.new
    wb = p.workbook

    wb.add_worksheet(:name => "Worksheet") do |sheet|
        test_data.each do |row_data|
            sheet.add_row row_data
        end
    end
    p.serialize('output/ruby.xlsx')

    end_timestamp = Time.now().to_f

    puts "Ruby: Writing 10000x50 cells of data takes #{end_timestamp - begin_timestamp} seconds"
end

test()
