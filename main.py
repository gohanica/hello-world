'''
Created on 2019/10/15

@author: K.Yamada
'''
import numpy as np
import matplotlib.pyplot as pit

FEATURES=10
LOOPTIMES=20

with open("data2019.txt",encoding="utf-8") as f:
    lst=f.readlines()
    print(lst)
    
    articleLst=[]
    
    for article in lst:
        article=article.replace("'","").replace("\"","").replace("-","").replace(".","").replace(",","").replace("?","").replace("(","").replace(")","")
        article=article.lower()
        articleLst.append(article)
    print(articleLst)
    
    word_set=set()
    tmp=[]
    
    for stri in articleLst:
        tmp=stri.split()
        word_set=word_set.union(set(tmp))
    print(word_set)
    
    LOOPTIMES=[]
    i=0
    while i<len(word_set):
        LOOPTIMES.append(0)
        i+=1
    
    word_dic=dict(zip(word_set,LOOPTIMES))
    print(word_dic)
    
    for stri in articleLst:
        for word in stri.split():
            if word in word_set:
                word_dic[word]+=1
    print(word_dic)
    
    onlyOneLst=[]
    
    for word,LOOPTIMES in word_dic.items():
        if LOOPTIMES<=1:
            onlyOneLst.append(word)
    for word in onlyOneLst:
        word_dic.pop(word)
        
    #irrelevantLst=["to","in","and","a","of","the","at","but","is","that","into","this","on","an","are","so","by","its","were","was"]
    #for word in irrelevantLst:
    #    word_dic.pop(word)
       
    print(word_dic)
    
    new_word_set=[]
    
    for word in word_dic:
        new_word_set.append(word)
    lstLst=[]
    for article in articleLst:
        zeroLst=[0]*len(new_word_set)
        article_word_dict=dict(zip(new_word_set,zeroLst))
        for word in article.split():
            keyLst=[]
            if word in new_word_set:
                article_word_dict[word]+=1
                
        for word,value in article_word_dict.items():
            keyLst.append(value)
            
        print(article_word_dict)
        
        lstLst.append(keyLst)
        print(lstLst)
        #行を結合
    articleArr=np.array(lstLst)
    print(articleArr)
    print(np.shape(articleArr))
    #課題2
    V=articleArr
    
    W=np.random.rand(len(articleLst),FEATURES)
    print(W)
    H=np.random.rand(FEATURES,len(new_word_set))
    print(H)
    
    i=0    
        
    while i<LOOPTIMES:
        for a in range(FEATURES):
            for m in range(len(new_word_set)):
                H_deno=np.dot(np.dot(W.T,W),H)[a,m]
                H_nume=np.dot(W.T,V)[a,m]
                H[a,m]=H[a,m]*(H_nume/H_deno)
        
        for j in range(len(articleLst)):
            for a in range(FEATURES):
                W_deno=np.dot(np.dot(W,H),H.T)[j,a]
                W_nume=np.dot(V,H.T)[j,a]
                W[j,a]=W[j,a]*(W_nume/W_deno)
    
        f_cost=0
    
        WH=np.dot(W,H)
        for a in range(len(articleLst)):
            for b in range(len(new_word_set)):
                f_cost+=(V[a,b]-WH[a,b])**2
        print(f_cost)
        i+=1
        pit.plot(i,f_cost,color='b',ms=2,lw=10,ls='-',alpha=0.9,marker='o')
            
    for p in range(FEATURES):
        W_lst=W[:,p].tolist()
        feat_W=dict(zip(articleLst,W_lst))    
        t=1
        for article,weight in sorted(feat_W.items(),key=lambda x:-x[1]):
        
            if t>3:
                break
            print(str(article)+":"+str(weight))
            t+=1
    
        print(W_lst)
    
        H_lst=H[p].tolist()
        print(H_lst)
        feat_H=dict(zip(new_word_set,H_lst))
        t=1
        for word,value in sorted(feat_H.items(),key=lambda x:-x[1]):
        
            if t>6:
                break
            print(str(word)+":"+str(value))
            t+=1
        
    #課題3
    pit.show()
    