﻿using System;
using System.Text;

namespace TyphenApi
{
    public abstract class TypeBase
    {
        static readonly JSONSerializer jsonSerializer = new JSONSerializer();
        static readonly QueryStringSerializer queryStringSerializer = new QueryStringSerializer();

        public override string ToString()
        {
            return ToJSON();
        }

        public static T FromJSON<T>(string text) where T : TypeBase, new()
        {
            return jsonSerializer.Deserialize<T>(Encoding.UTF8.GetBytes(text));
        }

        public string ToJSON()
        {
            return Encoding.UTF8.GetString(jsonSerializer.Serialize(this));
        }

        public string ToQueryString()
        {
            return Encoding.UTF8.GetString(queryStringSerializer.Serialize(this));
        }
    }
}
