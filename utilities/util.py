import pandas as pd
from sqlalchemy import create_engine

def excel_to_postgres(excel_file, table_name, db_connection_string):
    # Read Excel file
    df = pd.read_excel(excel_file)

    # Handle null values with pandas fillna
    # You can customize the replacement values based on column types
    replacements = {
        'numeric_columns': 0,        # for numeric columns
        'string_columns': '',        # for text columns
        'date_columns': None,        # keeps NULL for dates
        'boolean_columns': False     # for boolean columns
    }

    # Replace NaN values based on column type
    for col in df.columns:
        if pd.api.types.is_numeric_dtype(df[col]):
            df[col] = df[col].fillna(replacements['numeric_columns'])
        elif pd.api.types.is_string_dtype(df[col]):
            df[col] = df[col].fillna(replacements['string_columns'])
        elif pd.api.types.is_datetime64_dtype(df[col]):
            df[col] = df[col].fillna(replacements['date_columns'])
        elif pd.api.types.is_bool_dtype(df[col]):
            df[col] = df[col].fillna(replacements['boolean_columns'])

    # Create database connection
    engine = create_engine(db_connection_string)

    # Write to PostgreSQL
    df.to_sql(
        name=table_name,
        con=engine,
        if_exists='replace',  # or 'append' if you want to add to existing table
        index=False,
        chunksize=1000  # Process in chunks for large files
    )

# Usage example
connection_string = "postgresql://luho:luho@localhost:5432/myapp"
excel_to_postgres('./florence-venues.xlsx', 'florence_open_data', connection_string)
