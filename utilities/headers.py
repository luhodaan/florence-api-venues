
import pandas as pd

# Load the Excel file
df = pd.read_excel('utilities/florence-venues.xlsx')

# Print the headers (column names)
print(df.columns)