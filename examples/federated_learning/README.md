## virtual environment (recommended)
1. create 

    ```
    python -m venv venv
    ```

2. apply

    - for windows
    ```
    .\venv\Scripts\activate
    ```

    - for linux
    ```
    source ./venv/bin/activate
    ```

## install package
```
pip install --upgrade pip
pip install -r requirements.txt
```

## for client
1. download model from **DataShare**
    here we assume that 'mnist.py' is what we get

2. download initial parameters from **DataShare** (maybe not need)
    here we assume that 'results/avg/parameters' is what we get

3. train with local dataset
    ```
    python client.py
    ```
4. upload result to **DataShare**
    just simply upload output files in **results/client/**

## for server
1. upload model to **DataShare**
    omit..

2. waiting for private dataset provider training

3. download results from different private dataset providers
    here we assume that 'results/' are what we get

4. aggregate these results
    ```
    python server.py
    ```
    then we will get what we wanted in **results/avg/**
